package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	tokens []string
	result int
}{
	{
		[]string{"2", "1", "+", "3", "*"},
		9,
	},
	{
		[]string{"4", "13", "5", "/", "+"},
		6,
	},
	{
		[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
		22,
	},
	{
		[]string{"2", "2", "+", "2", "*"},
		8,
	},
	{
		[]string{"2", "2", "2", "*", "+"},
		6,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing/%d", i), func(t *testing.T) {
			output := evalRPN(testCase.tokens)
			if output != testCase.result {
				t.Errorf("Wrong answer. Expected %v, got %v", testCase.result, output)
			}
		})
	}
}
