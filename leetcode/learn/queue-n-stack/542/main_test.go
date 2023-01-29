package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	mat    [][]int
	expect [][]int
}{
	{
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
	},
	{
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 1, 1},
		},
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 2, 1},
		},
	},
	{
		[][]int{
			{1, 1, 0},
			{1, 1, 0},
			{1, 1, 0},
		},
		[][]int{
			{2, 1, 0},
			{2, 1, 0},
			{2, 1, 0},
		},
	},
	{
		[][]int{
			{1, 1, 0},
			{1, 1, 0},
			{1, 1, 1},
		},
		[][]int{
			{2, 1, 0},
			{2, 1, 0},
			{3, 2, 1},
		},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := updateMatrix(testCase.mat)

			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Wrong answer. Expect %v, got %v", testCase.expect, output)
			}
		})
	}
}
