package main

import (
	"math"
	"sort"
)

type Heap struct {
	data     []int
	size     int
	capacity int
}

func Construct(capacity int) Heap {
	data := make([]int, capacity)
	for i := 0; i < capacity; i++ {
		data[i] = math.MaxInt64
	}

	return Heap{
		data:     data,
		size:     0,
		capacity: capacity,
	}
}

func (h *Heap) Left(i int) int {
	return 2*i + 1
}

func (h *Heap) Right(i int) int {
	return 2*i + 2
}

func (h *Heap) Parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) Add(x int) {
	if h.size == h.capacity {
		return
	}

	h.size++
	i := h.size - 1
	h.data[i] = x

	for i != 0 && h.data[h.Parent(i)] > h.data[i] {
		h.data[h.Parent(i)], h.data[i] = h.data[i], h.data[h.Parent(i)]
		i = h.Parent(i)
	}
}

func (h *Heap) Heapify(i int) {
	smallest := i
	left := h.Left(i)
	right := h.Right(i)

	if left < len(h.data) && h.data[smallest] > h.data[left] {
		smallest = left
	}

	if right < len(h.data) && h.data[smallest] > h.data[right] {
		smallest = right
	}

	if smallest != i {
		h.data[smallest], h.data[i] = h.data[i], h.data[smallest]
		h.Heapify(smallest)
	}
}

func (h *Heap) Extract() {
	if h.size <= 0 {
		return
	}

	h.data[0] = h.data[h.size-1]
	h.data[h.size-1] = math.MaxInt64
	h.size--

	h.Heapify(0)
}

func (h *Heap) GetMin() int {
	return h.data[0]
}

func (h *Heap) GetSize() int {
	return h.size
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	heap := Construct(len(intervals))
	heap.Add(intervals[0][1])
	for i := 1; i < len(intervals); i++ {
		meeting := intervals[i]
		minEnding := heap.GetMin()

		if minEnding <= meeting[0] {
			heap.Extract()
		}

		heap.Add(meeting[1])
	}

	return heap.GetSize()
}
