package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	arr1   []int
	arr2   []int
	d      int
	expect int
}{
	{
		[]int{2, 1, 100, 3},
		[]int{-5, -2, 10, -3, 7},
		6,
		1,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := findTheDistanceValue(testCase.arr1, testCase.arr2, testCase.d)
			if output != testCase.expect {
				t.Errorf("Output %d is not equal to the expected value %d", output, testCase.expect)
			}
		})
	}
}
