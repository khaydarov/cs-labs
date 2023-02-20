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
		[]int{1, 2, 2, 1},
		[]int{2, 2},
		[]int{2, 2},
	},
	{
		[]int{1, 2, 2, 1, 4},
		[]int{2, 2},
		[]int{2, 2},
	},
	{
		[]int{1, 2, 2, 1, 4, 4},
		[]int{2, 2, 4},
		[]int{2, 2, 4},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := intersect(testCase.nums1, testCase.nums2)

			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, output)
			}
		})
	}
}
