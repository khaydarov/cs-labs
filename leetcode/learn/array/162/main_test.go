package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums     []int
	expected int
}{
	{
		[]int{1, 2, 3, 4},
		3,
	},
	{
		[]int{1, 2, 3, 1},
		2,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := findPeakElement(testCase.nums)
			if output != testCase.expected {
				t.Errorf("Output %d not equal to expected %d", output, testCase.expected)
			}
		})
	}
}
