package main

import (
	"math"
)

type HeapNode struct {
	Row    int
	Column int
	Value  int
}

type MinHeap struct {
	size     int
	capacity int
	data     []HeapNode
}

func (h *MinHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) left(i int) int {
	return 2*i + 1
}

func (h *MinHeap) right(i int) int {
	return 2*i + 2
}

func (h *MinHeap) heapify(i int) {
	l := h.left(i)
	r := h.right(i)

	smallest := i
	if l < h.size && h.data[l].Value < h.data[smallest].Value {
		smallest = l
	}

	if r < h.size && h.data[r].Value < h.data[smallest].Value {
		smallest = r
	}

	if i != smallest {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.heapify(smallest)
	}
}

func (h *MinHeap) Size() int {
	return h.size
}

func (h *MinHeap) Insert(x HeapNode) {
	if h.size == h.capacity {
		return
	}

	h.size++
	i := h.size - 1
	h.data[i] = x

	for i != 0 && h.data[h.parent(i)].Value > h.data[i].Value {
		h.data[h.parent(i)], h.data[i] = h.data[i], h.data[h.parent(i)]
		i = h.parent(i)
	}
}

func (h *MinHeap) GetMin() HeapNode {
	if h.Size() == 0 {
		return HeapNode{}
	}

	return h.data[0]
}

func (h *MinHeap) Extract() {
	if h.Size() == 0 {
		return
	}

	h.data[0] = h.data[h.size-1]
	h.data[h.size-1] = HeapNode{
		Value: math.MaxInt32,
	}
	h.size--

	h.heapify(0)
}

// TC: O(NLogN)
// SC: O(N)
func kthSmallest(matrix [][]int, k int) int {
	heapSize := len(matrix) // min(, k)
	heap := MinHeap{
		size:     0,
		capacity: heapSize,
		data:     make([]HeapNode, heapSize),
	}

	for i := 0; i < len(matrix); i++ {
		node := HeapNode{
			Row:    i,
			Column: 0,
			Value:  matrix[i][0],
		}

		heap.Insert(node)
	}

	for k > 1 {
		k--
		v := heap.GetMin()
		heap.Extract()

		if v.Column < len(matrix[v.Row])-1 {
			heap.Insert(HeapNode{
				Row:    v.Row,
				Column: v.Column + 1,
				Value:  matrix[v.Row][v.Column+1],
			})
		}
	}

	return heap.GetMin().Value
}
