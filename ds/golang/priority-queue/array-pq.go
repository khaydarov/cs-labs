package priority_queue

func ConstructArrayPriorityQueue(capacity int) ArrayPriorityQueue {
	return ArrayPriorityQueue{
		capacity: capacity,
	}
}

type ArrayPriorityQueue struct {
	items    []Item
	capacity int
	size     int
}

func (q *ArrayPriorityQueue) Front() int {
	if q.Size() == 0 {
		return -1
	}

	return q.items[0].value
}

func (q *ArrayPriorityQueue) Size() int {
	return q.size
}

func (q *ArrayPriorityQueue) Insert(value, priority int) bool {
	if q.isFull() {
		return false
	}

	newItem := NewItem(value, priority)

	var added bool
	for position, item := range q.items {
		if priority < item.priority {
			q.insertIntoPosition(position, newItem)
			added = true
			break
		}
	}

	if !added {
		q.insertIntoPosition(len(q.items), newItem)
	}

	return true
}
func (q *ArrayPriorityQueue) insertIntoPosition(position int, item Item) {
	rest := make([]Item, len(q.items[position:]))
	copy(rest, q.items[position:])

	q.items = append(q.items[:position], item)
	q.items = append(q.items, rest...)
}

func (q *ArrayPriorityQueue) Dequeue() bool {
	return false
}

func (q *ArrayPriorityQueue) isFull() bool {
	return q.size == q.capacity
}
