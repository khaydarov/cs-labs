package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums   []int
	k      int
	expect int64
}{
	{
		[]int{4, 1, 5, 3, 1},
		3,
		12,
	},
	{
		[]int{4, 6, 2},
		3,
		12,
	},
	{
		[]int{1, 3, 5},
		1,
		-1,
	},
	{
		[]int{2, 1, 1},
		2,
		2,
	},
	{
		[]int{5, 7, 9, 3, 11},
		5,
		-1,
	},
	{
		[]int{20, 20, 3, 1, 12, 10, 3, 7, 6, 6},
		9,
		82,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := largestEvenSum(testCase.nums, testCase.k)
			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, output)
			}
		})
	}
}
