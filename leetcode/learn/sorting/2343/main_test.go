package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	nums    []string
	queries [][]int
	result  []int
}{
	{
		nums:    []string{"3", "6", "9", "1"},
		queries: [][]int{{1, 1}},
		result:  []int{3},
	},
	{
		nums:    []string{"102", "473", "251", "814"},
		queries: [][]int{{1, 1}, {2, 3}, {4, 2}, {1, 2}},
		result:  []int{2, 2, 1, 0},
	},
	{
		nums:    []string{"24", "37", "96", "04"},
		queries: [][]int{{2, 1}, {2, 2}},
		result:  []int{3, 0},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			result := SmallestTrimmedNumbers(testCase.nums, testCase.queries)
			if !reflect.DeepEqual(result, testCase.result) {
				t.Errorf("expected %v, got %v", testCase.result, result)
			}
		})
	}
}
