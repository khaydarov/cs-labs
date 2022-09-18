package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums1  []int
	nums2  []int
	result float64
}{
	{
		[]int{1, 3},
		[]int{2},
		2,
	},
	{
		[]int{1, 2},
		[]int{3, 4},
		2.5,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := findMedianSortedArrays(testCase.nums1, testCase.nums2)
			if output != testCase.result {
				t.Errorf("Output %f not equal to expected %f", output, testCase.result)
			}
		})
	}
}
