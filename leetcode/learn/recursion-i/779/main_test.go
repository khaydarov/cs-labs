package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	n      int
	k      int
	expect int
}{
	{
		1,
		1,
		0,
	},
	{
		2,
		1,
		0,
	},
	{
		2,
		2,
		1,
	},
	{
		4,
		4,
		0,
	},
	{
		4,
		5,
		1,
	},
	{
		4,
		7,
		0,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := kthGrammar(testCase.n, testCase.k)

			if testCase.expect != output {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, output)
			}
		})
	}
}
