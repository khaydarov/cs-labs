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
		[]int{2, 7, 11, 15},
		9,
		[]int{1, 2},
	},
	{
		[]int{2, 7, 11, 15},
		26,
		[]int{3, 4},
	},
	{
		[]int{2, 3, 4},
		6,
		[]int{1, 3},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := twoSum(testCase.nums, testCase.target)

			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, output)
			}
		})
	}
}
