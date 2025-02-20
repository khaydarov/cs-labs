package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	given    int
	expected int
}{
	{
		given:    123,
		expected: 321,
	},
	{
		given:    -123,
		expected: -321,
	},
	{
		given:    120,
		expected: 21,
	},
	{
		given:    0,
		expected: 0,
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			actual := reverse(tc.given)
			if actual != tc.expected {
				t.Errorf("expected: %d, actual: %d", tc.expected, actual)
			}
		})
	}
}
