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
		[]int{1, 3, 4, 2, 2},
		2,
	},
	{
		[]int{3, 1, 3, 4, 2},
		3,
	},
	{
		[]int{1, 1},
		1,
	},
	{
		[]int{1, 1, 2},
		1,
	},
	{
		[]int{1, 2, 2},
		2,
	},
	{
		[]int{2, 2, 2, 2, 2},
		2,
	},
	{
		[]int{1, 3, 4, 2, 1},
		1,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := findDuplicate(testCase.nums)
			if output != testCase.expect {
				t.Errorf("Not equal to expected %d", testCase.expect)
			}
		})
	}

}
