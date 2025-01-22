package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums     []int
	expected int
}{
	{
		nums:     []int{3, 6, 9, 1},
		expected: 3,
	},
	{
		nums:     []int{10},
		expected: 0,
	},
	{
		nums:     []int{1, 10000000},
		expected: 9999999,
	},
	{
		nums:     []int{1, 1, 1, 1},
		expected: 0,
	},
	{
		nums:     []int{1, 1000, 2000},
		expected: 1000,
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			actual := maximumGap(tc.nums)
			if actual != tc.expected {
				t.Fatalf("Expected %v but got %v", tc.expected, actual)
			}
		})
	}
}
