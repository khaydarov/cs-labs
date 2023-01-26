package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	s      string
	expect string
}{
	{
		"3[a]2[bc]",
		"aaabcbc",
	},
	{
		"3[a2[c]]",
		"accaccacc",
	},
	{
		"2[abc]3[cd]ef",
		"abcabccdcdcdef",
	},
	{
		"abc3[cd]xyz",
		"abccdcdcdxyz",
	},
	{
		"11[x]c",
		"xxxxxxxxxxxc",
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := decodeString(testCase.s)
			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, output)
			}
		})
	}
}
