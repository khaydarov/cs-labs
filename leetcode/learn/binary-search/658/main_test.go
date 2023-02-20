package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	arr      []int
	k        int
	x        int
	expected []int
}{
	{
		[]int{1, 1, 2, 2, 2, 2, 2, 3, 3},
		3,
		3,
		[]int{2, 3, 3},
	},
	{
		[]int{1, 2, 3, 4, 4, 4, 4, 5, 5},
		3,
		3,
		[]int{2, 3, 4},
	},
	{
		[]int{1, 2, 3, 4, 5},
		3,
		5,
		[]int{3, 4, 5},
	},
	{
		[]int{1, 2, 3, 4, 5},
		3,
		4,
		[]int{3, 4, 5},
	},
	{
		[]int{-2, -1, 1, 2, 3, 4, 5},
		7,
		3,
		[]int{-2, -1, 1, 2, 3, 4, 5},
	},
	{
		[]int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8},
		2,
		2,
		[]int{1, 3},
	},
	{
		[]int{1, 1, 1, 10, 10, 10},
		1,
		9,
		[]int{10},
	},
	{
		[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8},
		3,
		5,
		[]int{3, 3, 4},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := findClosestElements(testCase.arr, testCase.k, testCase.x)

			if !reflect.DeepEqual(output, testCase.expected) {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expected, output)
			}
		})
	}
}
