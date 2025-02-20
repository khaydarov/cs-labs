package main

import "testing"

var testCases = []struct {
	num      uint32
	expected uint32
}{
	{43261596, 964176192},
	{4294967293, 3221225471},
	{0, 0},
}

func TestSolution1(t *testing.T) {
	for _, tc := range testCases {
		got := reverseBits1(tc.num)
		if got != tc.expected {
			t.Fatalf("expected %v but got %v", tc.expected, got)
		}
	}
}

func TestSolution2(t *testing.T) {
	for _, tc := range testCases {
		got := reverseBits2(tc.num)
		if got != tc.expected {
			t.Fatalf("expected %v but got %v", tc.expected, got)
		}
	}
}

func TestSolution3(t *testing.T) {
	for _, tc := range testCases {
		got := reverseBits3(tc.num)
		if got != tc.expected {
			t.Fatalf("expected %v but got %v", tc.expected, got)
		}
	}
}
