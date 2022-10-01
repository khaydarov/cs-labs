package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	matrix [][]int
	expect int
}{
	{
		[][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}},
		7,
	},
	{
		[][]int{{0, 1}, {1, 1}},
		3,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := countSquares(testCase.matrix)
			if output != testCase.expect {
				t.Errorf("Not equal to expected %d", testCase.expect)
			}
		})
	}

}
