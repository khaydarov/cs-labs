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
		[]int{1, 1, 1},
		1,
	},
	{
		[]int{1, 2, 2, 2, 5, 10},
		6,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := waysToSplit(testCase.nums)
			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, output)
			}
		})
	}
}
