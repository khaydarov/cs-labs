package heap

import (
	"fmt"
	"testing"
)

func TestHeapMinValue(t *testing.T) {
	testCases := []struct {
		capacity  int
		nums      []int
		expectMin int
	}{
		{
			5,
			[]int{16, 9, 8, 18, 3},
			3,
		},
		{
			3,
			[]int{1, 2, 3, 4},
			1,
		},
		{
			1,
			[]int{1, 2, 3, 4},
			1,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testcase:%d", i), func(t *testing.T) {
			h := Construct(testCase.capacity)
			for _, v := range testCase.nums {
				h.Insert(v)
			}

			output := h.Get()
			if output != testCase.expectMin {
				t.Errorf("Heap min value is not correct, expected %d, got %d", testCase.expectMin, output)
			}
		})
	}
}

func TestHeapExtract(t *testing.T) {
	testCases := []struct {
		capacity     int
		nums         []int
		extractCount int
		expectMin    int
	}{
		{
			5,
			[]int{16, 9, 8, 18, 3},
			1,
			8,
		},
		{
			3,
			[]int{1, 2, 3, 4},
			2,
			3,
		},
		{
			1,
			[]int{1, 2, 3, 4},
			1,
			-1,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testcase:%d", i), func(t *testing.T) {
			h := Construct(testCase.capacity)
			for _, v := range testCase.nums {
				h.Insert(v)
			}

			for j := 0; j < testCase.extractCount; j++ {
				h.Extract()
			}

			output := h.Get()
			if output != testCase.expectMin {
				t.Errorf("Heap min value is not correct, expected %d, got %d", testCase.expectMin, output)
			}
		})
	}
}
