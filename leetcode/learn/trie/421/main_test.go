package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums []int
	want int
}{
	{
		nums: []int{3, 10, 5, 25, 2, 8},
		want: 28,
	},
	{
		nums: []int{0},
		want: 0,
	},
	{
		nums: []int{3, 10, 5, 25, 2, 8, 24},
		want: 29,
	},
}

func TestSolution1(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			got := findMaximumXOR1(tc.nums)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			got := findMaximumXOR2(tc.nums)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
