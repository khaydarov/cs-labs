package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums   []int
	target int
	expect int
}{
	{
		[]int{2, 3, 1, 2, 4, 3},
		7,
		2,
	},
	{
		[]int{1, 4, 4},
		1,
		1,
	},
	{
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		11,
		0,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := minSubArrayLen(testCase.target, testCase.nums)
			if output != testCase.expect {
				t.Errorf("Not equal to expected %d", testCase.expect)
			}
		})
	}
}
