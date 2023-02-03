package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	x        int
	expected int
}{
	{
		4,
		2,
	},
	{
		9,
		3,
	},
	{
		5,
		2,
	},
	{
		1,
		1,
	},
	{
		25,
		5,
	},
	{
		26,
		5,
	},
	{
		35,
		5,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test.%d", i), func(t *testing.T) {
			output := mySqrt1(testCase.x)

			if output != testCase.expected {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expected, output)
			}
		})
	}
}
