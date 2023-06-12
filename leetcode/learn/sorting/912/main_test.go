package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	nums   []int
	expect []int
}{
	{
		[]int{2, 0, 2, 1, 1, 0},
		[]int{0, 0, 1, 1, 2, 2},
	},
	{
		[]int{2, 0, 1},
		[]int{0, 1, 2},
	},
	{
		[]int{1, 2, 0},
		[]int{0, 1, 2},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			sortArray(testCase.nums)

			if !reflect.DeepEqual(testCase.nums, testCase.expect) {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, testCase.nums)
			}
		})
	}
}
