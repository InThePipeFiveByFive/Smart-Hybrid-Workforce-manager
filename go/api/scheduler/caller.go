package scheduler

import (
	"bytes"
	"context"
	"encoding/json"
	"lib/clock"
	"lib/restclient"
	"net/http"
	"os"
	"time"
)

var (
	HTTPClient  restclient.HTTPClient
	Clock       clock.Clock   = &clock.RealClock{} // TODO: @JonathanEnslin make sure this is not a bad way of doing it
	timeout     time.Duration = 1800 * time.Second
	endpointURL string
)

func init() {
	// get env vars here
	endpointURL = os.Getenv("SCHEDULER_ADDR")

	HTTPClient = &http.Client{
		Timeout: timeout,
	}
}

var DaysOfWeek = map[string]time.Weekday{
	"Sunday":    time.Sunday,
	"Monday":    time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Saturday":  time.Saturday,
}

// datesEqual checks the equality of dates but ignoring time
func datesEqual(t1 time.Time, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.YearDay() == t2.YearDay()
}

// mayCall determines using the passed in params, and the current time, whether or not the scheduler
// may be called
func mayCall(scheduledDay string, now time.Time) bool {
	// check if correct day of week
	return scheduledDay == now.Weekday().String()
}

// TimeOfNextWeekDay returns the date/time of the next 'weekday'
func TimeOfNextWeekDay(now time.Time, weekday string) time.Time {
	day := int(DaysOfWeek[weekday])
	currentDay := int(now.Weekday())
	daysUntil := int((day - currentDay + 7) % 7) // +7 Is to ensure that the firs part of the expr is always >= 0
	if daysUntil == 0 {
		daysUntil = 7
	}
	y, m, d := now.AddDate(0, 0, daysUntil).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Location())
}

// CheckAndCall will access the logs, and then call the scheduler if the info log
// entry permits it
func checkAndCall(now time.Time, scheduledDay string) error {
	// scheduledDay := "Friday" // TODO: @JonathanEnslin Make env var
	if mayCall(scheduledDay, now) {
		// TODO: @JonathanEnslin move this into a seperate function that uses exponential backoff
		nextMonday := TimeOfNextWeekDay(now, "Monday")            // Start of next week
		nextSaturday := TimeOfNextWeekDay(nextMonday, "Saturday") // End of next work-week
		schedulerData, err := GetSchedulerData(nextMonday, nextSaturday)
		if err != nil {
			return err
		}
		err = Call(schedulerData) // TODO: @JonathanEnslin handle the return data
		if err != nil {
			return err
		}
	}
	return nil
}

// Call Calls the scheduler endpoint and passes it the data
func Call(data *SchedulerData) error { // TODO: @JonathanEnslin improve this function
	body, _ := json.Marshal(data)
	bodyBytesBuff := bytes.NewBuffer(body)

	request, err := http.NewRequest(http.MethodPost, endpointURL, bodyBytesBuff)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	response, err := HTTPClient.Do(request) // TODO: @JonathanEnslin URL env param
	if err != nil {
		return err
	}
	defer response.Body.Close()
	candidateBookings := &CandidateBookings{}
	err = json.NewDecoder(response.Body).Decode(candidateBookings)
	if err != nil {
		return err
	}
	err = makeBookings(*candidateBookings)
	if err != nil {
		return err
	}
	return nil
}

// callOnDay will call checkAndCall() on each recurring certain day of the week,
// the method can be cancelled using the passed in context
func callOnDay(ctx context.Context, scheduledDay string) {
	// Initial call, for when the function initially gets called
	_ = checkAndCall(time.Now(), scheduledDay)
	// periodic calls
	stopLoop := false
	for !stopLoop {
		nextDay := TimeOfNextWeekDay(time.Now(), scheduledDay) // TODO: @JonathanEnslin, allow scheduled day to be changed
		timer := time.NewTimer(time.Until(nextDay))
		defer timer.Stop()
		select {
		case <-timer.C:
			_ = checkAndCall(time.Now(), scheduledDay)
		case <-ctx.Done():
			stopLoop = true
		}
	}
}
