package main

import "fmt"

const MaxHeap = 1
const MinHeap = -1

type Heap struct {
	data 		[]int
	capacity 	int
	size		int
	mode 		int
}

func (h *Heap) Parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) Left(i int) int {
	return 2 * i + 1
}

func (h *Heap) Right(i int) int {
	return 2 * i + 2
}

func (h *Heap) GetMin() int {
	return h.data[0]
}

// Insert adds new value to the heap
// TC: O(log N)
// SC: O(1)
func (h *Heap) Insert(k int) {
	if h.size == h.capacity {
		return
	}

	h.size++

	i := h.size - 1
	h.data[i] = k

	if h.mode == MinHeap {
		for i != 0 && h.data[h.Parent(i)] > h.data[i] {
			h.data[i], h.data[h.Parent(i)] = h.data[h.Parent(i)], h.data[i]
			i = h.Parent(i)
		}
	} else {
		for i != 0 && h.data[h.Parent(i)] < h.data[i] {
			h.data[i], h.data[h.Parent(i)] = h.data[h.Parent(i)], h.data[i]
			i = h.Parent(i)
		}
	}
}

func (h *Heap) Extract() int {
	if h.size <= 0 {
		return -1
	}

	if h.size == 0 {
		h.size--
		return h.data[0]
	}

	root := h.data[0]
	h.data[0] = h.data[h.size - 1]
	h.data[h.size - 1] = 0
	h.size--

	if h.mode == MinHeap {
		h.MinHeapify(0)
	} else {
		h.MaxHeapify(0)
	}
	
	return root
}

// MinHeapify adjusts heap
func (h *Heap) MinHeapify(i int) {
	left := h.Left(i)
	right := h.Right(i)

	smallest := i
	if left < h.size && h.data[left] < h.data[smallest] {
		smallest = left
	}

	if right < h.size && h.data[right] < h.data[smallest] {
		smallest = right
	}

	if smallest != i {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.MinHeapify(smallest)
	}
}

func (h *Heap) MaxHeapify(i int) {
	left := h.Left(i)
	right := h.Right(i)

	biggest := i
	if left < h.size && h.data[left] > h.data[biggest] {
		biggest = left
	}

	if right < h.size && h.data[right] > h.data[biggest] {
		biggest = right
	}

	if biggest != i {
		h.data[biggest], h.data[i] = h.data[i], h.data[biggest]
		h.MaxHeapify(biggest)
	}
}

func main() {
	nums := make([]int, 5)
	heap := Heap{data: nums, capacity: 5, mode: MaxHeap}

	heap.Insert(10)
	heap.Insert(35)
	heap.Insert(3)
	heap.Insert(12)
	fmt.Println(heap)
	fmt.Println(heap.Extract())
	fmt.Println(heap)
}