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
		[]int{1, 2, 3, 4, 7, 7, 7, 7, 7, 7},
		1,
	},
	{
		[]int{1, 2, 3, 4, 7, 7, 7, 7, 7, 7, 0},
		0,
	},
	{
		[]int{1, 3, 5},
		1,
	},
	{
		[]int{2, 2, 2, 2, 0, 1},
		0,
	},
	{
		[]int{4, 5, 5, 5, 5, 0, 1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4},
		0,
	},
	{
		[]int{3, 3, 1, 3},
		1,
	},
	{
		[]int{1, 1, 1},
		1,
	},
	{
		[]int{10, 1, 10, 10, 10},
		1,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := findMin(testCase.nums)
			if output != testCase.expected {
				t.Errorf("Output %d not equal to expected %d", output, testCase.expected)
			}
		})
	}
}
