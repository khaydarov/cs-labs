package main

import "fmt"

type MinHeap struct {
	data []int
	size int
	cap  int
}

func (h *MinHeap) Left(i int) int {
	return 2 * i + 1
}

func (h *MinHeap) Right(i int) int {
	return 2 * i + 2
}

func (h *MinHeap) Parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) Insert(x int) {
	if h.size == h.cap {
		return
	}

	h.size++
	i := h.size - 1
	h.data = append(h.data, x)

	for i != 0 && h.data[h.Parent(i)] > h.data[i] {
		h.data[h.Parent(i)], h.data[i] = h.data[i], h.data[h.Parent(i)]
		i = h.Parent(i)
	}
}

func (h *MinHeap) GetMin() int {
	return h.data[0]
}

func (h *MinHeap) Extract() {
	if h.size == 0 {
		return
	}

	i := h.size - 1
	h.data[0] = h.data[i]
	h.size--

	h.Heapify(0)
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

	if smallest != i {
		h.data[smallest], h.data[i] = h.data[i], h.data[smallest]
		h.Heapify(smallest)
	}
}

func (h *MinHeap) ReplaceMin(x int) {
	h.data[0] = x
	h.Heapify(0)
}

func (h *MinHeap) Empty() bool {
	return h.size == 0
}

type KthLargest struct {
	heap 	*MinHeap
	k 		int
}


func Constructor(k int, nums []int) KthLargest {
	heap := &MinHeap{
		cap: k,
	}

	if len(nums) > 0 {
		for i := 0; i < k; i++ {
			if i < len(nums) {
				heap.Insert(nums[i])
			}
		}

		for i := k; i < len(nums); i++ {
			if nums[i] > heap.GetMin() {
				heap.ReplaceMin(nums[i])
			}
		}
	}

	return KthLargest{
		heap,
		k,
	}
}

func (this *KthLargest) Add(val int) int {
	if this.heap.Empty() || this.k > this.heap.size {
		this.heap.Insert(val)
	} else if val > this.heap.GetMin() {
		this.heap.ReplaceMin(val)
	}

	return this.heap.GetMin()
}

func main() {
	nums := []int{}
	k := Constructor(1, nums)
	fmt.Println(k.Add(-1), k.heap.data)
	fmt.Println(k.Add(1), k.heap.data)
	fmt.Println(k.Add(-2), k.heap.data)
	fmt.Println(k.Add(-4), k.heap.data)
	fmt.Println(k.Add(3), k.heap.data)
}