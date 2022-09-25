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
			{1, 2, 3},
			{6, 7, 8},
			{11, 12, 13},
			{16, 17, 18},
		},
		[]int{1, 2, 6, 11, 7, 3, 8, 12, 16, 17, 13, 18},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := findDiagonalOrder(testCase.matrix)
			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Not equal to expected %d", testCase.expect)
			}
		})
	}
}
