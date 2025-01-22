package radix

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	nums     []int
	expected []int
}{
	{
		nums:     []int{170, 45, 75, 90, 802, 24, 2, 66},
		expected: []int{2, 24, 45, 66, 75, 90, 170, 802},
	},
	{
		nums:     []int{170, 45, 75, 90, 802, 24, 2, 66, 1000},
		expected: []int{2, 24, 45, 66, 75, 90, 170, 802, 1000},
	},
	{
		nums:     []int{170, 45, 75, 90, 802, 24, 2, 66, 1000, 1001},
		expected: []int{2, 24, 45, 66, 75, 90, 170, 802, 1000, 1001},
	},
	{
		nums:     []int{1, 2, 3, 0, -1, -2},
		expected: []int{-2, -1, 0, 1, 2, 3},
	},
}

func TestRadisSort(t *testing.T) {
	for _, tc := range testCases {
		RadixSort(tc.nums)
		if !reflect.DeepEqual(tc.nums, tc.expected) {
			t.Fatalf("Expected %v but got %v", tc.expected, tc.nums)
		}
	}
}
