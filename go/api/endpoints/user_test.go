package endpoints

import (
	"api/data"
	"encoding/json"
	"io/ioutil"
	dtdb "lib/dockertest_db"
	tu "lib/testutils"
	ts "lib/test_setup"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Taken from lib/utils, copied since it is not exported
type errorResponse struct {
	Error map[string]interface{} `json:"error"`
}

func createUser(identifier string, firstName string, lastName string, email string, picture string) data.User {
	var user data.User
	user.Identifier = &identifier
	user.FirstName = &firstName
	user.LastName = &lastName
	user.Email = &email
	user.Picture = &picture
	return user
}

func TestRegisterUserHandler(t *testing.T) {
	// todo @JonathanEnslin find out if no password/firstname/lastname should be allowed
	testdb := ts.SetupTest(t)
	defer dtdb.StopTestDbWithTest(testdb, t, false)

	// ==================
	// Perform tests ====
	// ==================
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	type badExpect struct {
		responseCode    int
		responseMessage string
	}

	// Basic Bad tests ================
	basicBadTests := []struct {
		name    string
		request string
		args    args
		expect  badExpect
	}{
		{
			name: "Bad JSON, syntax error",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/booking/information`, strings.NewReader(`
				{
					first_name: "John",
					last_name: "Smith",
					email: "johnsmith@test.com",
					picture: "/pic.jpg"
					password : "password"
				}`)),
			},
			expect: badExpect{
				responseCode:    http.StatusBadRequest,
				responseMessage: "invalid_request",
			},
		},
		{
			name: "Bad JSON, array",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/booking/information`, strings.NewReader(`
				[
					{
						"first_name": "anonymous",
						"last_name": "anonymous",
						"email": "asd@test.com",
						"password" : "password"
					}
				]`)),
			},
			expect: badExpect{
				responseCode:    http.StatusBadRequest,
				responseMessage: "invalid_request",
			},
		},
		{
			name: "Bad Email", // Regex might have to be tested more extensively
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/booking/information`, strings.NewReader(`
				{
					"first_name": "anonymous",
					"last_name": "anonymous",
					"email": "asd@@test.com",
					"password" : "password"
				}`)),
			},
			expect: badExpect{
				responseCode:    http.StatusBadRequest,
				responseMessage: "invalid_email",
			},
		},
	}

	for _, tt := range basicBadTests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterUserHandler(tt.args.w, tt.args.r) // Make request
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
			}
			defer response.Body.Close()
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Error(tu.Scolourf(tu.RED, "Could not read response body, err: %v", err))
				t.FailNow()
			}
			var _error errorResponse
			err = json.Unmarshal(bodyBytes, &_error) // decode body
			if err != nil {
				t.Errorf(tu.Scolourf(tu.RED, "Invalid JSON, could not decode, err: %v", err))
				t.FailNow()
			}
			if message, ok := _error.Error["message"]; ok {
				if message != tt.expect.responseMessage {
					t.Error(tu.Scolourf(tu.RED, "Incorrect message returned, expected '%s', got '%s'", tt.expect.responseMessage, message))
				}
			} else {
				t.Error(tu.Scolourf(tu.RED, "Expected an error message, got none"))
			}
		})
	}

	type goodExpect struct {
		responseCode int
		responseBody *string
		user         data.User
	}

	// ==================
	// Good tests ================
	requestBodies := make([]*string, 1) // len should match len(goodTests) todo @JonathanEnslin update to use constructor later, eliminate need for this

	goodTests := []struct {
		name    string
		request string
		args    args
		expect  goodExpect
	}{
		{
			name: "OK registration",
			args: args{
				httptest.NewRecorder(),
				httptest.NewRequest(http.MethodPost, `http://localhost:8100/api/booking/information`, strings.NewReader(`
					{
						"first_name": "John",
						"last_name": "Smith",
						"email": "johnsmith@test.com",
						"picture": "/picture.png",
						"password" : "password"
					}`)),
			},
			expect: goodExpect{
				responseCode: http.StatusOK,
				responseBody: requestBodies[0], // do not use yet, API returns null
				user:         createUser("johnsmith@test.com", "John", "Smith", "johnsmith@test.com", "/picture.png"),
			},
		},
	}

	for _, tt := range goodTests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterUserHandler(tt.args.w, tt.args.r) // Make request
			// check response code
			response := tt.args.w.Result()
			if response.StatusCode != tt.expect.responseCode {
				t.Error(tu.Scolourf(tu.RED, "Invalid response code recieved, expected %d, got %d", tt.expect.responseCode, response.StatusCode))
				t.FailNow()
			}
			defer response.Body.Close()
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Error(tu.Scolourf(tu.RED, "Could not read response body, err: %v", err))
				t.FailNow()
			}
			var responseString *string
			err = json.Unmarshal(bodyBytes, &responseString) // decode body
			if err != nil {
				t.Errorf(tu.Scolourf(tu.RED, "Invalid response, could not decode, err: %v", err))
				t.FailNow()
			}
			if (responseString == nil && tt.expect.responseBody != nil) || (responseString != nil && tt.expect.responseBody == nil) {
				t.Error(tu.Scolourf(tu.RED, "Response incorrect, expected: %v address, got %v address", tt.expect.responseBody, responseString))
			} else if responseString != nil && *responseString != *tt.expect.responseBody {
				t.Error(tu.Scolourf(tu.RED, "Response incorrect, expected: %v, got %v", *tt.expect.responseBody, *responseString))
			}

			rows, err := testdb.Db.Query(`SELECT COUNT(*) FROM "user".identifier;`)
			if err != nil {
				t.Error(tu.Scolourf(tu.RED, "Could not query db, err: %v", err))
				t.FailNow()
			}
			rows.Next() // assumes if query was succesful there will be one row
			var numRows int
			rows.Scan(&numRows)
			if numRows != 1 {
				t.Error(tu.Scolourf(tu.RED, "Expected number of rows after insertion to be 1, there are: %d", numRows))
				t.FailNow()
			}
			// Check if DB was updated correctly
			// NB, assumption is made that there was no mock data inserted at this point, db is reset after each loop
			rows, err = testdb.Db.Query(`SELECT id, identifier, first_name, last_name, email, picture, date_created FROM "user".identifier;`)
			if err != nil {
				t.Error(tu.Scolourf(tu.RED, "Could not query db, err: %v", err))
				t.FailNow()
			}
			var user data.User
			rows.Next()
			err = rows.Scan(
				&user.Id,
				&user.Identifier,
				&user.FirstName,
				&user.LastName,
				&user.Email,
				&user.Picture,
				&user.DateCreated,
			)
			if err != nil {
				t.Error(tu.Scolourf(tu.RED, "Could not scan rows into user, err: %v", err))
				t.FailNow()
			}
			// no null check for identifier since db already checks this
			if *user.Identifier != *tt.expect.user.Identifier {
				t.Error(tu.Scolourf(tu.RED, "Identifer queried from DB does not equal expected, got %v, expected %v", *user.Identifier, *tt.expect.user.Identifier))
			}

			// todo @JonathanEnslin finish
		})
	}
}
