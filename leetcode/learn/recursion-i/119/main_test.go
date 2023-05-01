package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	rowIndex int
	expected []int
}{
	{
		0,
		[]int{1},
	},
	{
		1,
		[]int{1, 1},
	},
	{
		2,
		[]int{1, 2, 1},
	},
	{
		3,
		[]int{1, 3, 3, 1},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := getRow(testCase.rowIndex)

			if !reflect.DeepEqual(output, testCase.expected) {
				t.Errorf("Wrong answer. Expected %v, got %v", testCase.expected, output)
			}
		})
	}
}
