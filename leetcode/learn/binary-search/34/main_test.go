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
		[]int{5, 7, 7, 8, 8, 10},
		6,
		[]int{-1, -1},
	},
	{
		[]int{5, 7, 7, 8, 8, 10},
		8,
		[]int{3, 4},
	},
	{
		[]int{5, 7, 7, 8, 8, 10},
		5,
		[]int{0, 0},
	},
	{
		[]int{5, 7, 7, 8, 8, 10},
		10,
		[]int{5, 5},
	},
	{
		[]int{1, 1, 1, 1, 1},
		1,
		[]int{0, 4},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := searchRange(testCase.nums, testCase.target)

			fmt.Print(testCase, output)
			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, output)
			}
		})
	}
}
