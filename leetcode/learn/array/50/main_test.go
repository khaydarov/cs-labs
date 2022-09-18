package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	x      float64
	n      int
	expect float64
}{
	{
		2,
		3,
		8,
	},
	{
		3,
		2,
		9,
	},
	{
		1,
		10,
		1,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := myPow(testCase.x, testCase.n)
			if output != testCase.expect {
				t.Errorf("Output %f not equal to expected %f", output, testCase.expect)
			}
		})
	}
}
