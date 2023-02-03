package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums   []int
	expect int
}{
	{
		[]int{1, 2, 3, 4, 5},
		1,
	},
	{
		[]int{5, 0, 1, 2, 3, 4},
		0,
	},
	{
		[]int{4, 5, 0, 1, 2, 3},
		0,
	},
	{
		[]int{3, 4, 5, 0, 1, 2},
		0,
	},
	{
		[]int{2, 3, 4, 5, 0, 1},
		0,
	},
	{
		[]int{1, 2, 3, 4, 5, 0},
		0,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := findMin(testCase.nums)

			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, output)
			}
		})
	}
}
