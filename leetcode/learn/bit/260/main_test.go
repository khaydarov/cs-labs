package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	nums []int
	want []int
}{
	{
		nums: []int{2, 2, 1, 1, 3, 4},
		want: []int{3, 4},
	},
	{
		nums: []int{1, 2, 1, 3, 2, 5},
		want: []int{3, 5},
	},
	{
		nums: []int{1, 2, 1, 3, 2, 5, 5, 6},
		want: []int{3, 6},
	},
	{
		nums: []int{1, 2, 1, -3, 2, 5, 5, 6, 6, -1},
		want: []int{-1, -3},
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		got := singleNumber(tc.nums)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("failed for %+v: got=%+v, want=%+v", tc.nums, got, tc.want)
		}
	}
}
