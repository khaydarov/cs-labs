package priority_queue

type Item struct {
	value    int
	priority int
}

func NewItem(value, priority int) Item {
	return Item{value, priority}
}

// PriorityQueue is a Data Structure contract
type PriorityQueue interface {
	Front() int
	Size() int
	Insert() bool
	Dequeue() bool
}
