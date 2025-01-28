package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	num      uint32
	expected uint32
}{
	{
		num:      5,
		expected: 2684354560,
	},
	{
		num:      43261596,
		expected: 964176192,
	},
	{
		num:      4294967293,
		expected: 3221225471,
	},
	{
		num:      1,
		expected: 2147483648,
	},
	{
		num:      0,
		expected: 0,
	},
	{
		num:      10,
		expected: 1342177280,
	},
	{
		num:      100,
		expected: 637534208,
	},
}

func TestSolution1(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			actual := reverseBits1(tc.num)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			actual := reverseBits2(tc.num)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
