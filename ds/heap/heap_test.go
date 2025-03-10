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

func TestNativeHeap(t *testing.T) {
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
			h := Construct(10)
			for _, v := range testCase.nums {
				h.Add(v)
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

func TestNativeHeapBuild(t *testing.T) {
	testCases := []struct {
		nums     []int
		minValue int
	}{
		{
			nums:     []int{6, 4, 1, 8, 9},
			minValue: 1,
		},
		{
			nums:     []int{1, 2, 3},
			minValue: 1,
		},
		{
			nums:     []int{99, 101, 55},
			minValue: 55,
		},
		{
			nums:     []int{99, 101, 55, -1, 0, 9, 17, -16},
			minValue: -16,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testcase:%d", i), func(t *testing.T) {
			h := Build(testCase.nums)

			output := h.Get()
			if output != testCase.minValue {
				t.Errorf("Heap min value is not correct, expected %d, got %d", testCase.minValue, output)
			}
		})
	}
}
