package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	input    [][]int
	expected [][]int
}{
	{
		[][]int{
			{2147483647, -1, 0, 2147483647},
			{2147483647, 2147483647, 2147483647, -1},
			{2147483647, -1, 2147483647, -1},
			{0, -1, 2147483647, 2147483647},
		},
		[][]int{
			{3, -1, 0, 1},
			{2, 2, 1, -1},
			{1, -1, 2, -1},
			{0, -1, 3, 4},
		},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			wallsAndGates(testCase.input)
			if !reflect.DeepEqual(testCase.input, testCase.expected) {
				t.Errorf("Wrong answer. Expected %v, but got %v", testCase.expected, testCase.input)
			}
		})
	}
}
