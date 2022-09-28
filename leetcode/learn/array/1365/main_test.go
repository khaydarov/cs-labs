package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	nums   []int
	expect []int
}{
	{
		[]int{8, 1, 2, 2, 3},
		[]int{4, 0, 1, 1, 3},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := smallerNumbersThanCurrent(testCase.nums)
			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Not equal to expected %d", testCase.expect)
			}
		})
	}
}
