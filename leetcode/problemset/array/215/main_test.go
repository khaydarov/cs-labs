package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums   []int
	k      int
	expect int
}{
	{
		[]int{3, 2, 1, 5, 6, 4},
		2,
		5,
	},
	{
		[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		4,
		4,
	},
	{
		[]int{-10, 10, 20, 30},
		2,
		20,
	},
	{
		[]int{-10, 10, 20, 30},
		3,
		10,
	},
	{
		[]int{-10, 10, 20, 30},
		4,
		-10,
	},
	{
		[]int{-10, -5, -2, 0, 10, 20, 30},
		4,
		0,
	},
	{
		[]int{-10, -5, -2, 0, 10, 20, 30},
		5,
		-2,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := findKthLargest(testCase.nums, testCase.k)
			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, output)
			}
		})
	}
}
