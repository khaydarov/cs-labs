package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	str string
	num int
}{
	{"M", 1000},
	{"CM", 900},
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
	{"MCDDD", 2400},
	{"XIV", 14},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			got := romanToInt(tc.str)
			if got != tc.num {
				t.Errorf("expected %d, got %d", tc.num, got)
			}
		})
	}
}
