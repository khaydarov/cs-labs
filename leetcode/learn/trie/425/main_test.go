package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	words    []string
	expected [][]string
}{
	{
		words: []string{"abat", "baba", "atan", "atal"},
		expected: [][]string{
			{"baba", "abat", "baba", "atan"},
			{"baba", "abat", "baba", "atal"},
		},
	},
	{
		words: []string{"area", "lead", "wall", "lady", "ball"},
		expected: [][]string{
			{"wall", "area", "lead", "lady"},
			{"ball", "area", "lead", "lady"},
		},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			actual := wordSquares(tc.words)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
