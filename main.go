package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)

	fmt.Println(*h)
	val := old[n-1]
	*h = old[:n-1]

	return val
}

func connectSticks(sticks []int) int {
	minHeap := MinHeap{}
	for _, v := range sticks {
		heap.Push(&minHeap, v)
	}

	fmt.Println(heap.Pop(&minHeap))
	// cost := 0
	// for len(minHeap) > 1 {
	// 	x := heap.Pop(&minHeap)
	// 	y := heap.Pop(&minHeap)

	// 	s := x.(int) + y.(int)
	// 	cost += s
	// 	heap.Push(&minHeap, s)
	// }

	return 1
}

func main() {
	connectSticks([]int{2, 4, 1})
	// fmt.Println(r)
}
