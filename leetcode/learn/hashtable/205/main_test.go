package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	s      string
	t      string
	expect bool
}{
	{
		"egg",
		"add",
		true,
	},
	{
		"foo",
		"bar",
		false,
	},
	{
		"paper",
		"title",
		true,
	},
	{
		"badc",
		"baba",
		false,
	},
	{
		"bbdc",
		"baba",
		false,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			output := isIsomorphic(testCase.s, testCase.t)
			if output != testCase.expect {
				t.Errorf(
					"Wrong answer on s = %v, t = %v, output = %v, expect = %v",
					testCase.s,
					testCase.t,
					output,
					testCase.expect,
				)
			}
		})
	}
}
