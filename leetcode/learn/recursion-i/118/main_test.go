package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	numRows  int
	expected [][]int
}{
	{
		1,
		[][]int{
			{1},
		},
	},
	{
		2,
		[][]int{
			{1},
			{1, 1},
		},
	},
	{
		3,
		[][]int{
			{1},
			{1, 1},
			{1, 2, 1},
		},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := generate(testCase.numRows)

			if !reflect.DeepEqual(output, testCase.expected) {
				t.Errorf("Wrong answer. Expected %v, got %v", testCase.expected, output)
			}
		})
	}
}
