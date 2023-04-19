package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums     []int
	k        int
	expected int
}{
	{
		[]int{7, 2, 5, 10, 8},
		2,
		18,
	},
	{
		[]int{1, 2, 3, 4, 5},
		2,
		9,
	},
	{
		[]int{7, 2, 5, 10, 8},
		3,
		14,
	},
	{
		[]int{10, 5, 13, 4, 8, 4, 5, 11, 14, 9, 16, 10, 20, 8},
		8,
		25,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := splitArray(testCase.nums, testCase.k)

			if output != testCase.expected {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expected, output)
			}
		})
	}
}
