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
		[]int{1, 2, 3, 4, 5, 6},
		6,
		5,
	},
	{
		[]int{1, 2, 3, 4, 5, 6},
		3,
		2,
	},
	{
		[]int{1, 2, 3, 4, 5, 6},
		1,
		0,
	},
	{
		[]int{1, 2, 3, 4, 5, 6},
		9,
		-1,
	},
	{
		[]int{4, 5, 6, 0, 1, 2, 3},
		0,
		3,
	},
	{
		[]int{4, 5, 6, 0, 1, 2, 3},
		3,
		6,
	},
	{
		[]int{4, 5, 6, 0, 1, 2, 3},
		4,
		0,
	},
	{
		[]int{4, 5, 6, 7, 8, 9, 1, 2},
		8,
		4,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := search1(testCase.nums, testCase.target)

			if output != testCase.expected {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expected, output)
			}
		})
	}
}
