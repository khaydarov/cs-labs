package priority_queue

type PqItem struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type HeapPriorityQueue []*PqItem

func (pq HeapPriorityQueue) Len() int {
	return len(pq)
}

// We want Pop to give us the highest, not lowest, priority so we use greater than here.
func (pq HeapPriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq HeapPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *HeapPriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PqItem)
	item.index = n

	*pq = append(*pq, item)
}

func (pq *HeapPriorityQueue) Pop() any {
	old := *pq
	n := len(old)

	item := (*pq)[n]
	item.index = -1
	old[n-1] = nil

	*pq = old[:n-1]
	return item
}
