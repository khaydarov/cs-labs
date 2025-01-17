package main

import (
	"container/heap"
)

const (
	minHeap = iota
	maxHeap
)

type Heap struct {
	heapType int
	data     []int
}

func (h Heap) Less(i, j int) bool {
	if h.heapType == minHeap {
		return h.data[i] < h.data[j]
	}
	return h.data[i] > h.data[j]
}

func (h Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Push(x interface{}) {
	(*h).data = append((*h).data, x.(int))
}

func (h *Heap) Pop() interface{} {
	size := len((*h).data)
	val := (*h).data[size-1]
	(*h).data = (*h).data[:size-1]

	return val
}

type MedianFinder struct {
	lo *Heap
	hi *Heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		lo: &Heap{
			heapType: maxHeap,
			data:     make([]int, 0),
		},
		hi: &Heap{
			heapType: minHeap,
			data:     make([]int, 0),
		},
	}
}

// Time Complexity: O(log N)
// Space Complexity: O(1)
func (m *MedianFinder) AddNum(num int) {
	heap.Push(m.lo, num)
	if m.lo.Len() > 0 {
		max := heap.Pop(m.lo)
		heap.Push(m.hi, max)
	}

	if m.lo.Len() < m.hi.Len() {
		min := heap.Pop(m.hi)
		heap.Push(m.lo, min)
	}
}

// Time Complexity: O(1)
// Space Complexity: O(1)
func (m *MedianFinder) FindMedian() float64 {
	if m.lo.Len() > m.hi.Len() {
		return float64(m.lo.data[0])
	}
	return float64(m.lo.data[0]+m.hi.data[0]) * 0.5
}
