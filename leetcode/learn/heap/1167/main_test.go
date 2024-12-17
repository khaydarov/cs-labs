package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	input  []int
	output int
}{
	{
		input:  []int{2, 4, 3},
		output: 14,
	},
	{
		input:  []int{5},
		output: 0,
	},
	{
		input:  []int{1, 8, 3, 5},
		output: 30,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := connectSticks(testCase.input)
			if output != testCase.output {
				t.Errorf("Wrong answer. Expected %v, but got %v", testCase.output, output)
			}
		})
	}
}
