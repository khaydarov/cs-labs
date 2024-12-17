package heap

type MinIntHeap []int

func (h MinIntHeap) Len() int {
	return len(h)
}

func (h MinIntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() interface{} {
	heapDerefrenced := *h

	size := len(heapDerefrenced)
	val := heapDerefrenced[size-1]
	*h = heapDerefrenced[:size-1]

	return val
}
