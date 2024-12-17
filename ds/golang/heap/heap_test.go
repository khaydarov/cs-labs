package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeapMinValue(t *testing.T) {
	testCases := []struct {
		nums      []int
		expectMin int
	}{
		{
			[]int{16, 9, 8, 18, 3},
			3,
		},
		{
			[]int{1, 2, 3, 4},
			1,
		},
		{
			[]int{1, 2, 3, 4},
			1,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testcase:%d", i), func(t *testing.T) {
			h := MinIntHeap{}
			for _, v := range testCase.nums {
				heap.Push(&h, v)
			}

			output := heap.Pop(&h)
			if output != testCase.expectMin {
				t.Errorf("Heap min value is not correct, expected %d, got %d", testCase.expectMin, output)
			}
		})
	}
}

func TestHeapExtract(t *testing.T) {
	testCases := []struct {
		nums         []int
		extractCount int
		expectMin    int
	}{
		{
			[]int{16, 9, 8, 18, 3},
			1,
			8,
		},
		{
			[]int{1, 2, 3, 4},
			2,
			3,
		},
		{
			[]int{1, 2, 3, 4},
			1,
			2,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testcase:%d", i), func(t *testing.T) {
			h := MinIntHeap{}
			for _, v := range testCase.nums {
				heap.Push(&h, v)
			}

			for j := 0; j < testCase.extractCount; j++ {
				heap.Pop(&h)
			}

			output := heap.Pop(&h)
			if output != testCase.expectMin {
				t.Errorf("Heap min value is not correct, expected %d, got %d", testCase.expectMin, output)
			}
		})
	}
}
