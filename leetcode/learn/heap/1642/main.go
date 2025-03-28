package main

import (
	"container/heap"
)

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := &minHeap{}
	for i := 0; i < len(heights)-1; i++ {
		diff := heights[i+1] - heights[i]
		if diff <= 0 {
			continue
		}

		heap.Push(h, diff)
		if h.Len() > ladders {
			bricks -= heap.Pop(h).(int)
		}

		if bricks < 0 {
			return i
		}
	}

	return len(heights) - 1
}
