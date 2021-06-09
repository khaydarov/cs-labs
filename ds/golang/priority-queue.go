package main

import "fmt"

type QueueItem struct {
	Val 		int
	Priority 	int
}

type PriorityQueue struct {
	items []QueueItem
}

// Insert adds to the queue a new element with passed priority
//
// TC: O(N)
// SC: O(1)
func (p *PriorityQueue) Insert(val, priority int) {
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
func (p *PriorityQueue) Size() int {
	return len(p.items)
}

// Front returns value with lowest priority
func (p *PriorityQueue) Front() int {
	if len(p.items) == 0 {
		return -1
	}

	return p.items[0].Val
}

func main() {
	q := PriorityQueue{}
	q.Insert(1, 1)
	q.Insert(2, 2)
	q.Insert(4, 4)
	q.Insert(3, 3)
	q.Insert(5, 0)

	fmt.Println(q)
}