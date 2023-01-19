package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	deadends []string
	target   string
	expected int
}{
	{
		[]string{"0201", "0101", "0102", "1212", "2002"},
		"0202",
		6,
	},
	{
		[]string{"8888"},
		"0009",
		1,
	},
	{
		[]string{"0000"},
		"8888",
		-1,
	},
	{
		[]string{"0009"},
		"0009",
		-1,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := openLock(testCase.deadends, testCase.target)
			if output != testCase.expected {
				t.Errorf("Wrong answer. Expected %v, but got %v", testCase.expected, output)
			}
		})
	}
}
