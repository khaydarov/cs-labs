package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	text    string
	pattern string
	expect  bool
}{
	{
		"abbb",
		"ab*",
		true,
	},
	{
		"abc",
		"ab*",
		false,
	},
	{
		"abc",
		"ab*c",
		true,
	},
	{
		"abcd",
		"ab*c*.",
		true,
	},
	{
		"abaa",
		"a.*a*",
		true,
	},
	{
		"ab",
		".*c",
		false,
	},
	{
		"aaa",
		"a*a",
		true,
	},
	{
		"a",
		"ab*c*",
		true,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := IsMatch(testCase.text, testCase.pattern)

			if output != testCase.expect {
				t.Errorf("Expected %v but got %v", testCase.expect, output)
			}
		})
	}

}
