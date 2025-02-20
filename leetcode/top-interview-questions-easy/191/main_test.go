package main

import "testing"

var testCases = []struct {
	num      uint32
	expected int
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 1},
	{5, 2},
	{6, 2},
	{7, 3},
	{8, 1},
	{9, 2},
	{10, 2},
	{11, 3},
	{12, 2},
}

func TestSolution1(t *testing.T) {
	for _, tc := range testCases {
		actual := hammingWeight1(tc.num)
		if actual != tc.expected {
			t.Errorf("expected %d, got %d", tc.expected, actual)
		}
	}
}

func TestSolution2(t *testing.T) {
	for _, tc := range testCases {
		actual := hammingWeight2(tc.num)
		if actual != tc.expected {
			t.Errorf("expected %d, got %d", tc.expected, actual)
		}
	}
}
