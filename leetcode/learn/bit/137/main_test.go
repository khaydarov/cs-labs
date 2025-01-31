package main

import "testing"

var testCases = []struct {
	nums []int
	want int
}{
	{
		nums: []int{2, 2, 2, 1, 1, 1, -5},
		want: -5,
	},
	{
		nums: []int{1, 1, 1, 11, 22, 22, 22},
		want: 11,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		got := singleNumber(tc.nums)
		if got != tc.want {
			t.Errorf("nums: %v, got: %v, want: %v", tc.nums, got, tc.want)
		}
	}
}
