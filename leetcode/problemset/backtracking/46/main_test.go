package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	nums   []int
	expect [][]int
}{
	{
		[]int{1},
		[][]int{
			{1},
		},
	},
	{
		[]int{1, 2},
		[][]int{
			{1, 2},
			{2, 1},
		},
	},
	{
		[]int{1, 2, 3},
		[][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 2, 1},
			{3, 1, 2},
		},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := permute(testCase.nums)

			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, output)
			}
		})
	}
}
