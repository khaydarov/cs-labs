package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	a      int
	b      []int
	expect int
}{
	{
		2,
		[]int{1, 0},
		1024,
	},
	{
		2,
		[]int{3},
		8,
	},
	{
		1,
		[]int{4, 3, 3, 8, 5, 2},
		1,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := superPow(testCase.a, testCase.b)

			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, output)
			}
		})
	}
}
