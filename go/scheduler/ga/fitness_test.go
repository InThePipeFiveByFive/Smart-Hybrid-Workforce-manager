package ga

import (
	tu "lib/testutils"
	"reflect"
	"scheduler/data"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_distanceRadicand(t *testing.T) {
	type args struct {
		origin []float64
		coord  []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				origin: []float64{0.0, 0.0},
				coord:  []float64{4.0, 5.0},
			},
			want: 41.0,
		},
		{
			name: "Test 2",
			args: args{
				origin: []float64{1.0, 2.0},
				coord:  []float64{4.0, 6.0},
			},
			want: 25.0,
		},
		{
			name: "Test 3",
			args: args{
				origin: []float64{1.0, 2.5},
				coord:  []float64{-1.0, 6.0},
			},
			want: 16.25,
		},
		{
			name: "Test 4",
			args: args{
				origin: []float64{1.0},
				coord:  []float64{-1.0},
			},
			want: 4.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceRadicand(tt.args.origin, tt.args.coord); got != tt.want {
				t.Errorf("distanceRadicand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCentroid(t *testing.T) {
	type args struct {
		coords [][]float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Test 1",
			args: args{
				coords: [][]float64{{-1, -1}, {1, 3}, {0, 0}, {2, -1}, {2, 3}},
			},
			want: []float64{0.8, 0.8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCentroid(tt.args.coords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCentroid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_avgDistanceFromCentroid(t *testing.T) {
	type args struct {
		coords [][]float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				coords: [][]float64{{-1, -1}, {1, 3}, {0, 0}, {2, -1}, {2, 3}},
			},
			want: 2.1110702096228455,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := avgDistanceFromCentroid(tt.args.coords); got != tt.want {
				t.Errorf("avgDistanceFromCentroid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndividual_getTeamsGroupedByRooms(t *testing.T) {
	type args struct {
		domain *Domain
	}
	tests := []struct {
		name       string
		individual *Individual
		args       args
		want       []teamRoomGroups
	}{
		{
			name: "Test 1",
			args: args{
				domain: &Domain{
					SchedulerData: &data.SchedulerData{
						Teams: []*data.TeamInfo{
							{
								Team:    &data.Team{Id: tu.Ptr("Cabbage")},
								UserIds: []string{"Lime", "Lemon", "Grapefruit", "Banana"},
							},
							{
								Team:    &data.Team{Id: tu.Ptr("Broccoli")},
								UserIds: []string{"Blueberry", "Gooseberry", "Lemon"},
							},
							{
								Team:    &data.Team{Id: tu.Ptr("Lettuce")},
								UserIds: []string{"Strawberry", "Blueberry"},
							},
							{
								Team:    &data.Team{Id: tu.Ptr("Eggplant")},
								UserIds: []string{},
							},
						},
						Rooms: []*data.RoomInfo{
							{Room: &data.Room{Id: tu.Ptr("Freezer")}},
							{Room: &data.Room{Id: tu.Ptr("Fridge")}},
							{Room: &data.Room{Id: tu.Ptr("Pantry")}},
							{Room: &data.Room{Id: tu.Ptr("Countertop")}},
						},
						Resources: []*data.Resource{
							{Id: tu.Ptr("Shelf1"), RoomId: tu.Ptr("Freezer")},
							{Id: tu.Ptr("Shelf2"), RoomId: tu.Ptr("Freezer")},
							{Id: tu.Ptr("Shelf3"), RoomId: tu.Ptr("Freezer")},
							{Id: tu.Ptr("Shelf4"), RoomId: tu.Ptr("Freezer")},
							{Id: tu.Ptr("Shelf5"), RoomId: tu.Ptr("Freezer")},
							{Id: tu.Ptr("Shelf10"), RoomId: tu.Ptr("Fridge")},
							{Id: tu.Ptr("Shelf20"), RoomId: tu.Ptr("Fridge")},
							{Id: tu.Ptr("Shelf30"), RoomId: tu.Ptr("Fridge")},
							{Id: tu.Ptr("Shelf40"), RoomId: tu.Ptr("Fridge")},
							{Id: tu.Ptr("Shelf50"), RoomId: tu.Ptr("Fridge")},
							{Id: tu.Ptr("Shelf100"), RoomId: tu.Ptr("Pantry")},
							{Id: tu.Ptr("Shelf200"), RoomId: tu.Ptr("Pantry")},
							{Id: tu.Ptr("Shelf300"), RoomId: tu.Ptr("Pantry")},
							{Id: tu.Ptr("Shelf_1"), RoomId: tu.Ptr("Countertop")},
							{Id: tu.Ptr("Shelf_2"), RoomId: tu.Ptr("Countertop")},
							{Id: tu.Ptr("Shelf_3"), RoomId: tu.Ptr("Countertop")},
						},
					},
					InverseMap: map[string][]int{
						"Lime":       {0, 3},
						"Lemon":      {1},
						"Grapefruit": {4, 2},
						"Blueberry":  {5, 8},
						"Gooseberry": {6},
						"Strawberry": {7},
					},
				},
			},
			individual: &Individual{
				Gene: [][]string{
					//    0        1         2           3          4          5         6         7         8
					// Freezer  Freezer    Fridge     Pantry     Fridge     Freezer   Freezer   Freezer   Fridge
					{"Shelf1", "Shelf2", "Shelf10", "Shelf100", "Shelf30", "Shelf4", "Shelf5", "Shelf3", "Shelf20"},
				},
			},
			want: []teamRoomGroups{
				{
					teamId: "Broccoli",
					roomGroups: map[string][]int{
						"Freezer": {5, 6, 1},
						"Fridge":  {8},
					},
				},
				{
					teamId: "Cabbage",
					roomGroups: map[string][]int{
						"Freezer": {0, 1},
						"Fridge":  {2, 4},
						"Pantry":  {3},
					},
				},
				{
					teamId:     "Eggplant",
					roomGroups: map[string][]int{},
				},
				{
					teamId: "Lettuce",
					roomGroups: map[string][]int{
						"Freezer": {7, 5},
						"Fridge":  {8},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.domain.SchedulerData.ApplyMapping()
			got := tt.individual.getTeamsGroupedByRooms(tt.args.domain)
			assert.Len(t, got, len(tt.want), "Expected got to have a length of %v but got len %v", len(tt.want), len(got))
			sort.Slice(got, func(i, j int) bool {
				return got[i].teamId < got[j].teamId
			})
			for i := range tt.want {
				tu.MapsWithcSlicesMatchLoosely(t, got[i].roomGroups, tt.want[i].roomGroups)
			}
		})
	}
}
