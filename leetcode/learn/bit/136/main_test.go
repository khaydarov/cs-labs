package main

import "testing"

var testCases = []struct {
	nums []int
	want int
}{
	{
		nums: []int{2, 2, 1},
		want: 1,
	},
	{
		nums: []int{4, 1, 2, 1, 2},
		want: 4,
	},
	{
		nums: []int{1},
		want: 1,
	},
	{
		nums: []int{1, 1, 2},
		want: 2,
	},
	{
		nums: []int{1, 1, 2, 2, -3},
		want: -3,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		got := singleNumber(tc.nums)
		if got != tc.want {
			t.Errorf("nums: %v, want: %d, got: %d", tc.nums, tc.want, got)
		}
	}
}
