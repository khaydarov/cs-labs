package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums  []int
	pivot int
}{
	{
		[]int{1, 7, 3, 6, 5, 6},
		3,
	},
	{
		[]int{1, 2, 3},
		-1,
	},
	{
		[]int{2, 1, -1},
		0,
	},
	{
		[]int{1},
		0,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := pivotIndex(testCase.nums)
			if output != testCase.pivot {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.pivot, output)
			}
		})
	}
}
