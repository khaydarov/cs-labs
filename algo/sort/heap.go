package main

import (
	"fmt"
	"math"
)

type MinHeap struct {
	data []int
	size int
	cap  int
}

func (h *MinHeap) Left(i int) int {
	return 2 * i + 1
}

func (h *MinHeap) Parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) Right(i int) int {
	return 2 * i + 2
}

func (h *MinHeap) Add(x int) {
	if h.size == h.cap {
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

func (h *MinHeap) GetMin() int {
	if h.size == 0 {
		return math.MinInt32
	}

	return h.data[0]
}

func (h *MinHeap) Heapify(i int) {
	left := h.Left(i)
	right := h.Right(i)
	smallest := i

	if left < h.size && h.data[left] < h.data[smallest] {
		smallest = left
	}

	if right < h.size && h.data[right] < h.data[smallest] {
		smallest = right
	}

	if i != smallest {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.Heapify(smallest)
	}
}

func (h *MinHeap) ExtractMin() {
	if h.size == 0 {
		return
	}

	h.data[0] = h.data[h.size - 1]
	h.size--

	h.Heapify(0)
}

// HeapSort sorting algorithm implementation
//
// TC: O(N log N)
// SC: O(N log N)
func HeapSort(arr *[]int) {
	n := len(*arr)

	h := MinHeap{
		data: make([]int, n),
		cap: n,
	}

	for _, v := range *arr {
		h.Add(v)
	}

	for i := 0; i < n; i++ {
		(*arr)[i] = h.GetMin()
		h.ExtractMin()
	}
}

func main() {
	arr := []int{5, 4, 3, 1, 2, 9}
	HeapSort(&arr)
	fmt.Println(arr)
}