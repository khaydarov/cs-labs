package heap

import "math"

type MinHeap struct {
	size     int
	capacity int
	data     []int
}

func Construct(capacity int) MinHeap {
	data := make([]int, capacity)
	for i := 0; i < capacity; i++ {
		data[i] = math.MaxInt64
	}

	return MinHeap{
		capacity: capacity,
		data:     data,
	}
}

func Build(nums []int) MinHeap {
	minHeap := MinHeap{
		size:     len(nums),
		capacity: cap(nums),
		data:     nums,
	}

	startIdx := len(nums)/2 - 1
	for i := startIdx; i >= 0; i-- {
		minHeap.heapify(i)
	}

	return minHeap
}

func (h *MinHeap) Size() int {
	return h.size
}

func (h *MinHeap) Add(x int) {
	if h.size >= h.capacity {
		return
	}

	h.size++
	i := h.size - 1
	h.data[i] = x

	for i != 0 && h.data[h.parent(i)] > h.data[i] {
		h.data[h.parent(i)], h.data[i] = h.data[i], h.data[h.parent(i)]
		i = h.parent(i)
	}
}

func (h *MinHeap) Get() int {
	if h.size == 0 {
		return math.MaxInt64
	}
	return h.data[0]
}

func (h *MinHeap) Extract() {
	if h.size <= 0 {
		return
	}

	h.data[0] = h.data[h.size-1]
	h.data[h.size-1] = math.MaxInt64
	h.size--

	h.heapify(0)
}

func (h *MinHeap) heapify(i int) {
	smallest := i
	l := h.left(i)
	r := h.right(i)

	if l < h.size && h.data[l] < h.data[smallest] {
		smallest = l
	}

	if r < h.size && h.data[r] < h.data[smallest] {
		smallest = r
	}

	if smallest != i {
		h.data[smallest], h.data[i] = h.data[i], h.data[smallest]
		h.heapify(smallest)
	}
}

func (h MinHeap) left(i int) int {
	return 2*i + 1
}

func (h MinHeap) right(i int) int {
	return 2*i + 2
}

func (h MinHeap) parent(i int) int {
	return (i - 1) / 2
}
