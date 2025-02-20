package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	n    int
	want string
}{
	{1, "1"},
	{4, "1211"},
	{5, "111221"},
	{6, "312211"},
	{7, "13112221"},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			got := countAndSay(tc.n)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
