package bucket

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	nums     []int
	expected []int
}{
	{
		nums:     []int{3, 2, 1, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		nums:     []int{3, 1, 5, 4, 6},
		expected: []int{1, 3, 4, 5, 6},
	},
	{
		nums:     []int{-10, 3, 9, 8, -4, 0},
		expected: []int{-10, -4, 0, 3, 8, 9},
	},
	{
		nums:     []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		nums:     []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			BucketSort(testCase.nums)
			if !reflect.DeepEqual(testCase.nums, testCase.expected) {
				t.Fatalf("Expected %v but got %v", testCase.expected, testCase.nums)
			}
		})
	}
}
