package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	matrix [][]int
	expect []int
}{
	{
		[][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
		},
		[]int{1, 2, 3, 4, 5, 10, 15, 20, 19, 18, 17, 16, 11, 6, 7, 8, 9, 14, 13, 12},
	},
	{
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		[]int{1, 2, 3, 6, 9, 9, 7, 4, 5},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := SpiralCopy(testCase.matrix)
			if !reflect.DeepEqual(output, testCase.matrix) {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.matrix, output)
			}
		})
	}
}
