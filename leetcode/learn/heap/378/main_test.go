package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	matrix [][]int
	k      int
	expect int
}{
	{
		matrix: [][]int{
			{-1},
		},
		k:      1,
		expect: -1,
	},
	{
		matrix: [][]int{
			{2, 3},
			{3, 6},
			{1, 2},
		},
		k:      2,
		expect: 2,
	},
	{
		matrix: [][]int{
			{2, 3},
			{3, 6},
			{1, 2},
		},
		k:      3,
		expect: 2,
	},
	{
		matrix: [][]int{
			{1, 3, 5, 7, 9},
			{2, 4, 6, 8, 10},
			{11, 13, 15, 17, 19},
			{12, 14, 16, 18, 20},
			{21, 22, 23, 24, 25},
		},
		k:      8,
		expect: 8,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := kthSmallest(testCase.matrix, testCase.k)
			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Wrong answer. Expected %v, but got %v", testCase.expect, output)
			}
		})
	}
}
