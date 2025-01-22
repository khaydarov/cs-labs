package counting

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	array  []int
	expect []int
}{
	{
		array:  []int{5, 3, 4, 2},
		expect: []int{2, 3, 4, 5},
	},
	{
		array:  []int{5, 3, 4, 4, 1, 9, -3, -2},
		expect: []int{-3, -2, 1, 3, 4, 4, 5, 9},
	},
	{
		array:  []int{5, 3, 4, 4, 1, 9, -3, -2, 5, 3, 4, 4, 1, 9, -3, -2},
		expect: []int{-3, -3, -2, -2, 1, 1, 3, 3, 4, 4, 4, 4, 5, 5, 9, 9},
	},
	{
		array:  []int{1, 2, 3, 0, -2, -3},
		expect: []int{-3, -2, 0, 1, 2, 3},
	},
}

func TestCountingSort(t *testing.T) {
	for _, tc := range testCases {
		CountingSortWithShifting(tc.array)
		if !reflect.DeepEqual(tc.array, tc.expect) {
			t.Errorf("Expected %v but got %v", tc.expect, tc.array)
			break
		}
	}
}
