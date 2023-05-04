package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	n      int
	expect int
}{
	{
		1,
		1,
	},
	{
		2,
		2,
	},
	{
		3,
		5,
	},
	{
		4,
		14,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := numTrees(testCase.n)

			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, output)
			}
		})
	}
}
