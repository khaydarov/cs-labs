package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums   []int
	target int
	expect bool
}{
	{
		[]int{2, 5, 6, 0, 0, 1, 2},
		0,
		true,
	},
	{
		[]int{2, 5, 6, 0, 0, 1, 2},
		3,
		false,
	},
	{
		[]int{1, 0, 1, 1, 1},
		1,
		true,
	},
	{
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		13,
		true,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := search(testCase.nums, testCase.target)
			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, output)
			}
		})
	}
}
