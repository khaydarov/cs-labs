package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	n      int
	matrix [][]int
}{
	{
		1,
		[][]int{{1}},
	},
	{
		2,
		[][]int{{1, 2}, {4, 3}},
	},
	{
		3,
		[][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := generateMatrix(testCase.n)
			if !reflect.DeepEqual(output, testCase.matrix) {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.matrix, output)
			}
		})
	}
}
