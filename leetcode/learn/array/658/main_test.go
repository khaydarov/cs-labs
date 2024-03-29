package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	nums   []int
	k      int
	x      int
	expect []int
}{
	{
		[]int{1, 2, 5, 5, 6, 6, 7, 7, 8, 9},
		7,
		9,
		[]int{5, 6, 6, 7, 7, 8, 9},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := findClosestElements(testCase.nums, testCase.k, testCase.x)
			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Not equal to expected %d", testCase.expect)
			}
		})
	}
}
