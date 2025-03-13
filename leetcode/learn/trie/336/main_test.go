package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	input    []string
	expected [][]int
}{
	{
		input:    []string{"abcd", "dcba", "lls", "s", "sssll"},
		expected: [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}},
	},
	{
		input:    []string{"bat", "tab", "cat"},
		expected: [][]int{{0, 1}, {1, 0}},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			result := palindromePairs(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
