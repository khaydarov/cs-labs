package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	arr    []int
	expect int
}{
	{
		[]int{1, 4, 2, 5, 3},
		58,
	},
	{
		[]int{1, 2},
		3,
	},
	{
		[]int{10, 11, 12},
		66,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := sumOddLengthSubarrays(testCase.arr)
			if output != testCase.expect {
				t.Errorf("Output %d is not equal to expected %d", output, testCase.expect)
			}
		})
	}
}
