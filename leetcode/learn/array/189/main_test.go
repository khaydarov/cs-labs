package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	nums   []int
	k      int
	expect []int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		3,
		[]int{5, 6, 7, 1, 2, 3, 4},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			rotate(testCase.nums, testCase.k)
			if !reflect.DeepEqual(testCase.nums, testCase.expect) {
				t.Errorf("Not equal to expected %d", testCase.expect)
			}
		})
	}
}
