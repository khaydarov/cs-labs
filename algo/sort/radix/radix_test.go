package radix

import "testing"

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
}

func TestRadisSort(t *testing.T) {
	for _, tc := range testCases {
		RadixSort(tc.nums)
		for i := 0; i < len(tc.nums); i++ {
			if tc.nums[i] != tc.expected[i] {
				t.Fatalf("Expected %v but got %v", tc.expected, tc.nums)
			}
		}
	}
}
