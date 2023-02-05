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
		[]int{1, 3, 5},
		1,
	},
	{
		[]int{5, 1, 3},
		1,
	},
	{
		[]int{2, 2, 2, 0, 1},
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
