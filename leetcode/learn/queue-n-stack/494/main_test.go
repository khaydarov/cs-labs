package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums     []int
	target   int
	expected int
}{
	{
		[]int{1, 1, 1, 1, 1},
		3,
		5,
	},
	{
		[]int{1},
		1,
		1,
	},
	{
		[]int{1, 1, 1, 1, 3},
		3,
		6,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := findTargetSumWays(testCase.nums, testCase.target)

			if output != testCase.expected {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expected, output)
			}
		})
	}
}