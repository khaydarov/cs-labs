package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	s      string
	expect bool
}{
	{
		"aba",
		false,
	},
	{
		"a",
		false,
	},
	{
		"abab",
		true,
	},
	{
		"ababab",
		true,
	},
	{
		"abcabcabcabc",
		true,
	},
	{
		"abaababaab",
		true,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := repeatedSubstringPattern(testCase.s)
			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, output)
			}
		})
	}
}
