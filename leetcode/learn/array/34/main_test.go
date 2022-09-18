package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	nums   []int
	target int
	expect []int
}{
	{
		[]int{1, 2, 2, 2, 4},
		2,
		[]int{1, 3},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := searchRange(testCase.nums, testCase.target)

			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Output %d not equal to expected %d", output, testCase.expect)
			}
		})
	}
}
