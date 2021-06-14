package main

import (
	"fmt"
	"math"
)

type QueueItem struct {
	Val 		int
	Priority 	int
}

// PriorityQueueArray is structure for implementation of queue with priority
// Uses simple array order to store data
type PriorityQueueArray struct {
	items []QueueItem
}

// Insert adds to the queue a new element with passed priority
//
// TC: O(N)
// SC: O(1)
func (p *PriorityQueueArray) Insert(val, priority int) {
	if len(p.items) == 0 {
		p.items = append(p.items, QueueItem{val, priority})
	} else {
		added := false
		item := QueueItem{val, priority}
		for i, current := range p.items {
			if item.Priority < current.Priority {
				rest := make([]QueueItem, len(p.items[i:]))

				copy(rest, p.items[i:])
				p.items = append(p.items[:i], item)
				p.items = append(p.items, rest...)
				added = true
				break
			}
		}

		if !added {
			p.items = append(p.items, item)
		}
	}
}

// Size returns the queue size
func (p *PriorityQueueArray) Size() int {
	return len(p.items)
}

// Front returns value with lowest priority
func (p *PriorityQueueArray) Front() int {
	if len(p.items) == 0 {
		return -1
	}

	return p.items[0].Val
}

// ---------------------

// PriorityQueueHeap is a structure that implements queue with priority
// Uses heap data structure to store items
type PriorityQueueHeap struct {
	data []QueueItem
	size int
	cap  int
}

func (h *PriorityQueueHeap) Left(i int) int {
	return 2 * i + 1
}

func (h *PriorityQueueHeap) Parent(i int) int {
	return (i - 1) / 2
}

func (h *PriorityQueueHeap) Right(i int) int {
	return 2 * i + 2
}

// Insert adds new value with priority
//
// TC: O(N log N)
// SC: O(1)
func (h *PriorityQueueHeap) Insert(val int, priority int) {
	if h.size == h.cap {
		return
	}

	h.size++
	i := h.size - 1
	h.data[i] = QueueItem{val, priority}

	for i != 0 && h.data[h.Parent(i)].Priority > h.data[i].Priority {
		h.data[h.Parent(i)], h.data[i] = h.data[i], h.data[h.Parent(i)]
		i = h.Parent(i)
	}
}

func (h *PriorityQueueHeap) Front() int {
	if h.size == 0 {
		return math.MinInt32
	}

	return h.data[0].Val
}

func (h *PriorityQueueHeap) Heapify(i int) {
	left := h.Left(i)
	right := h.Right(i)
	smallest := i

	if left < h.size && h.data[left].Priority < h.data[smallest].Priority {
		smallest = left
	}

	if right < h.size && h.data[right].Priority < h.data[smallest].Priority {
		smallest = right
	}

	if i != smallest {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.Heapify(smallest)
	}
}

func main() {
	q := PriorityQueueArray{}
	q.Insert(1, 1)
	q.Insert(2, 2)
	q.Insert(4, 4)
	q.Insert(3, 3)
	q.Insert(5, 0)

	fmt.Println(q)
}