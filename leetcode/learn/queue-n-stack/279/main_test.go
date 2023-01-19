package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	num    int
	expect int
}{
	{
		12,
		3,
	},
	{
		13,
		2,
	},
	{
		9,
		1,
	},
	{
		1,
		1,
	},
	{
		10,
		2,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := numSquaresMath(testCase.num)

			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v, got %v", testCase.expect, output)
			}
		})
	}
}
