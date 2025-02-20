package main

import "testing"

var testCases = []struct {
	num      int
	expected bool
}{
	{27, true},
	{0, false},
	{9, true},
	{45, false},
	{1, true},
	{3, true},
	{2, false},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		actual := isPowerOfThree(tc.num)
		if actual != tc.expected {
			t.Errorf("expected: %v, got: %v", tc.expected, actual)
		}
	}
}
