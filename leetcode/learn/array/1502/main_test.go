package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	arr    []int
	expect bool
}{
	{
		[]int{3, 5, 1},
		true,
	},
	{
		[]int{1, 2, 4},
		false,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := canMakeArithmeticProgression(testCase.arr)
			if output != testCase.expect {
				t.Errorf(fmt.Sprintf("Output %v is not equals to expected %v", output, testCase.expect))
			}
		})
	}
}
