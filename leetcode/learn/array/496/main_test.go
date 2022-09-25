package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	nums1  []int
	nums2  []int
	expect []int
}{
	{
		[]int{4, 1, 2},
		[]int{1, 3, 4, 2},
		[]int{-1, 3, -1},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := nextGreaterElement(testCase.nums1, testCase.nums2)
			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Not equal to expected %d", testCase.expect)
			}
		})
	}
}
