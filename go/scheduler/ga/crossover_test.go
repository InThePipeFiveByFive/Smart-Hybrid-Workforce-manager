package ga

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoPointSwap(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
		xP1  int
		xP2  int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "Test1 filled arr",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP1:  2,
				xP2:  6,
			},
			want:  []int{-1, 1, 20, 30, 40, 50, 6, 7, 8, 9},
			want1: []int{-10, 10, 2, 3, 4, 5, 60, 70, 80, 90},
		},
		{
			name: "Test2 crossover point at end",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP1:  2,
				xP2:  10,
			},
			want:  []int{-1, 1, 20, 30, 40, 50, 60, 70, 80, 90},
			want1: []int{-10, 10, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Test2 crossover point at start",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP1:  0,
				xP2:  2,
			},
			want:  []int{-10, 10, 2, 3, 4, 5, 6, 7, 8, 9},
			want1: []int{-1, 1, 20, 30, 40, 50, 60, 70, 80, 90},
		},
		{
			name: "Test2 crossover point 1 > crossover point 2",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP1:  2,
				xP2:  1,
			},
			want:  []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want1: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := twoPointSwap(tt.args.arr1, tt.args.arr2, tt.args.xP1, tt.args.xP2)
			assert.True(t, reflect.DeepEqual(got, tt.want), "twoPointSwap() got = %v, want %v", got, tt.want)
			assert.True(t, reflect.DeepEqual(got1, tt.want1), "twoPointSwap() got = %v, want %v", got, tt.want)
		})
	}
}

func TestFindValid(t *testing.T) {
	type args struct {
		index          int
		parent         []string
		otherParentMap map[string]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Return mapped element",
			args: args{
				index:  0,
				parent: []string{"c", "f", "e", "b", "a", "m", "d", "g"},
				otherParentMap: map[string]int{
					"c": 2,
					"d": 3,
					"z": 4,
					"e": 5,
				},
			},
			want: "m",
		},
		{
			name: "Return element at index",
			args: args{
				index:  1,
				parent: []string{"c", "f", "e", "b", "a", "m", "d", "g"},
				otherParentMap: map[string]int{
					"c": 2,
					"d": 3,
					"z": 4,
					"e": 5,
				},
			},
			want: "f",
		},
		{
			name: "Return mapped element second test",
			args: args{
				index:  6,
				parent: []string{"c", "f", "e", "b", "a", "m", "d", "g"},
				otherParentMap: map[string]int{
					"c": 2,
					"d": 3,
					"z": 4,
					"e": 5,
				},
			},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindValid(tt.args.index, tt.args.parent, tt.args.otherParentMap)
			assert.Equal(t, tt.want, got, "FindValid() = %v, want %v", got, tt.want)
		})
	}
}

func TestPMX(t *testing.T) {
	type args struct {
		p1  []string
		p2  []string
		xP1 int
		xP2 int
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{
			name: "No mapping required test, normal crossover",
			args: args{
				p1:  []string{"a", "b", "c", "d", "e", "f", "g", "h", "j"},
				p2:  []string{"h", "j", "f", "e", "d", "c", "g", "a", "b"},
				xP1: 2,
				xP2: 5,
			},
			want:  []string{"a", "b", "f", "e", "d", "c", "g", "h", "j"},
			want1: []string{"h", "j", "c", "d", "e", "f", "g", "a", "b"},
		},
		{
			name: "No mapping required test, crossover entire chromosome",
			args: args{
				p1:  []string{"a", "b", "c", "d", "e", "f", "g", "h", "j"},
				p2:  []string{"h", "j", "f", "e", "d", "c", "g", "a", "b"},
				xP1: 0,
				xP2: 9,
			},
			want:  []string{"h", "j", "f", "e", "d", "c", "g", "a", "b"},
			want1: []string{"a", "b", "c", "d", "e", "f", "g", "h", "j"},
		},
		{
			name: "Mapping required crossover, arrays not permuations of eachother",
			args: args{
				p1:  []string{"a", "b", "c", "d", "z", "e", "f", "g"},
				p2:  []string{"c", "f", "e", "b", "a", "m", "d", "g"},
				xP1: 2,
				xP2: 6,
			},
			want:  []string{"z", "d", "e", "b", "a", "m", "f", "g"},
			want1: []string{"m", "f", "c", "d", "z", "e", "b", "g"},
		},
		{
			name: "Mapping required crossover, arrays are permuations of eachother",
			args: args{
				p1:  []string{"a", "b", "c", "d", "e", "f", "g"},
				p2:  []string{"c", "f", "e", "b", "a", "d", "g"},
				xP1: 2,
				xP2: 5,
			},
			want:  []string{"c", "d", "e", "b", "a", "f", "g"},
			want1: []string{"a", "f", "c", "d", "e", "b", "g"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PMX(tt.args.p1, tt.args.p2, tt.args.xP1, tt.args.xP2)
			assert.True(t, reflect.DeepEqual(got, tt.want), "PMX() got = %v, want %v", got, tt.want)
			assert.True(t, reflect.DeepEqual(got1, tt.want1), "PMX() got1 = %v, want %v", got1, tt.want1)
		})
	}
}

func TestPartiallyMappedFlattenCrossoverValid(t *testing.T) {
	type args struct {
		domain       *Domain
		individuals  Individuals
		numOffspring int
	}
	tests := []struct {
		name                string
		args                args
		errMessage          string
		originalParentGene1 [][]string
		originalParentGene2 [][]string
	}{
		{
			name: "Normal crossover (daily form)",
			args: args{
				domain: nil,
				individuals: Individuals{
					&Individual{
						Gene: [][]string{
							{"a", "b", "c", "d", "e", "f", "g"},
						},
						Fitness: 0.0,
					},
					&Individual{
						Gene: [][]string{
							{"c", "f", "e", "b", "a", "d", "g"},
						},
						Fitness: 0.0,
					},
				},
				numOffspring: 2,
			},
			originalParentGene1: [][]string{
				{"a", "b", "c", "d", "e", "f", "g"},
			},
			originalParentGene2: [][]string{
				{"c", "f", "e", "b", "a", "d", "g"},
			},
			errMessage: "Expceted parents and children to contain the same element",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PartiallyMappedFlattenCrossoverValid(tt.args.domain, tt.args.individuals, tt.args.numOffspring)
			parent1Gene, parent2Gene := tt.args.individuals[0].Gene, tt.args.individuals[1].Gene
			child1Gene, child2Gene := got[0].Gene, got[1].Gene
			assert.ElementsMatch(t, parent1Gene[0], child1Gene[0], tt.errMessage)
			assert.ElementsMatch(t, parent2Gene[0], child2Gene[0], tt.errMessage)
			assert.Len(t, got, tt.args.numOffspring, "Incorrect number of individuals returned, expected=%v, got=%v", tt.args.numOffspring, len(got))
			assert.True(t, reflect.DeepEqual(parent1Gene, tt.originalParentGene1), "Parent 1 gene has been changed")
			assert.True(t, reflect.DeepEqual(parent2Gene, tt.originalParentGene2), "Parent 2 gene has been changed")
		})
	}
}

func TestCrossoverCaller(t *testing.T) {

	mockCrossover := func(domain *Domain, individuals Individuals, numOffspring int) Individuals {
		return individuals[:numOffspring] // Mock crossover
	}

	mockSelectionOperator := func(domain *Domain, individuals Individuals, count int) Individuals {
		return individuals[:count] // Mock crossover
	}

	individuals := Individuals{
		&Individual{
			Fitness: 0.0,
			Gene:    [][]string{{"pear", "litchi"}},
		},
		&Individual{
			Fitness: 0.0,
			Gene:    [][]string{{"cabbage", "broccoli"}},
		},
		&Individual{
			Fitness: 0.0,
			Gene:    [][]string{{"peanut", "almond"}},
		},
	}

	type args struct {
		crossoverOperator CrossoverOperator
		domain            *Domain
		individuals       Individuals
		selectionFunc     Selection
		offspring         int
	}
	tests := []struct {
		name        string
		args        args
		expectedLen int
	}{
		{
			name: "Test for the correct number of individuals being returned, expect 0",
			args: args{
				crossoverOperator: mockCrossover,
				domain:            nil,
				selectionFunc:     mockSelectionOperator,
				offspring:         0,
				individuals:       individuals,
			},
			expectedLen: 0,
		},
		{
			name: "Test for the correct number of individuals being returned, expect 1",
			args: args{
				crossoverOperator: mockCrossover,
				domain:            nil,
				selectionFunc:     mockSelectionOperator,
				offspring:         1,
				individuals:       individuals,
			},
			expectedLen: 1,
		},
		{
			name: "Test for the correct number of individuals being returned, expect 2",
			args: args{
				crossoverOperator: mockCrossover,
				domain:            nil,
				selectionFunc:     mockSelectionOperator,
				offspring:         2,
				individuals:       individuals,
			},
			expectedLen: 2,
		},
		{
			name: "Test for the correct number of individuals being returned, exepct 100",
			args: args{
				crossoverOperator: mockCrossover,
				domain:            nil,
				selectionFunc:     mockSelectionOperator,
				offspring:         100,
				individuals:       individuals,
			},
			expectedLen: 100,
		},
		{
			name: "Test for the correct number of individuals being returned, expect 101",
			args: args{
				crossoverOperator: mockCrossover,
				domain:            nil,
				selectionFunc:     mockSelectionOperator,
				offspring:         101,
				individuals:       individuals,
			},
			expectedLen: 101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CrossoverCaller(
				tt.args.crossoverOperator,
				tt.args.domain, tt.args.individuals,
				tt.args.selectionFunc,
				tt.args.offspring,
			)
			assert.Len(t, got, tt.expectedLen, "Incorrect number of individuals returned, got=%v, expected=%v", len(got), tt.expectedLen)
		})
	}
}

func Test_OnePointCrossover(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
		xP   int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "OnePointCrossover crossover point beyond parent length",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP:   10,
			},
			want:  []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want1: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
		},
		{
			name: "OnePointCrossover crossover point smaller than 0",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP:   -20,
			},
			want:  []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
			want1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "OnePointCrossover crossover point valid",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP:   4,
			},
			want:  []int{-1, 1, 2, 3, 40, 50, 60, 70, 80, 90},
			want1: []int{-10, 10, 20, 30, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := onePointCrossover(tt.args.arr1, tt.args.arr2, tt.args.xP)
			assert.True(t, reflect.DeepEqual(got, tt.want), "onePointCrossover() got = %v, want %v", got, tt.want)
			assert.True(t, reflect.DeepEqual(got1, tt.want1), "onePointCrossover() got = %v, want %v", got, tt.want)
		})
	}
}

func Test_TwoPointCrossover(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
		xP1  int
		xP2  int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "TwoPointCrossover crossover point beyond parent length",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP1:  10,
				xP2:  10,
			},
			want:  []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want1: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
		},
		{
			name: "TwoPointCrossover crossover point smaller than 0",
			args: args{
				arr1: []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
				xP1:  -20,
				xP2:  -20,
			},
			want:  []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want1: []int{-10, 10, 20, 30, 40, 50, 60, 70, 80, 90},
		},
		{
			name: "TwoPointCrossover first crossover point smaller than 0 the other greateer than 0",
			args: args{
				arr1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				arr2: []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
				xP1:  -20,
				xP2:  5,
			},
			want:  []int{-1, -2, -3, -4, -5, 6, 7, 8, 9, 10},
			want1: []int{1, 2, 3, 4, 5, -6, -7, -8, -9, -10},
		},
		{
			name: "TwoPointCrossover crossover point valid",
			args: args{
				arr1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				arr2: []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
				xP1:  2,
				xP2:  4,
			},
			want:  []int{1, 2, -3, -4, 5, 6, 7, 8, 9, 10},
			want1: []int{-1, -2, 3, 4, -5, -6, -7, -8, -9, -10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := twoPointCrossover(tt.args.arr1, tt.args.arr2, tt.args.xP1, tt.args.xP2)
			assert.True(t, reflect.DeepEqual(got, tt.want), "twoPointCrossover() got = %v, want %v", got, tt.want)
			assert.True(t, reflect.DeepEqual(got1, tt.want1), "twoPointCrossover() got = %v, want %v", got, tt.want)
		})
	}
}

func Test_CycleCrossover(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "CycleCrossover crossover valid",
			args: args{
				arr1: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
				arr2: []int{8, 2, 6, 7, 1, 5, 4, 0, 3},
			},
			want:  []int{0, 2, 6, 3, 1, 5, 4, 7, 8},
			want1: []int{8, 1, 2, 7, 4, 5, 6, 0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := cycleCrossover(tt.args.arr1, tt.args.arr2)
			assert.True(t, reflect.DeepEqual(got, tt.want), "cycleCrossover() got = %v, want %v", got, tt.want)
			assert.True(t, reflect.DeepEqual(got1, tt.want1), "cycleCrossover() got = %v, want %v", got, tt.want)
		})
	}
}
