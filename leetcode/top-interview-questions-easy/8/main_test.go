package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	numStr string
	want   int
}{
	{"42", 42},
	{"   -42", -42},
	{"4193 with words", 4193},
	{"words and 987", 0},
	{"-91283472332", -2147483648},
	{"91283472332", 2147483647},
	{"", 0},
	{"     ", 0},
	{"- 3-3", 0},
	{"-2147483647", -2147483647},
	{"2147483648", 2147483647},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			got := myAtoi(tc.numStr)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
