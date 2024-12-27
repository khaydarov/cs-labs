package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	heights []int
	bricks  int
	ladders int
	answer  int
}{
	{
		heights: []int{4, 2, 7, 6, 9, 14, 12},
		bricks:  5,
		ladders: 1,
		answer:  4,
	},
	{
		heights: []int{4, 12, 2, 7, 3, 18, 20, 3, 19},
		bricks:  10,
		ladders: 2,
		answer:  7,
	},
	{
		heights: []int{14, 3, 19, 3},
		bricks:  17,
		ladders: 0,
		answer:  3,
	},
	{
		heights: []int{1, 1000},
		bricks:  100,
		ladders: 0,
		answer:  0,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := furthestBuilding(testCase.heights, testCase.bricks, testCase.ladders)

			if output != testCase.answer {
				t.Errorf("Wrong answer. Expected %v, but got %v", testCase.answer, output)
			}
		})
	}
}
