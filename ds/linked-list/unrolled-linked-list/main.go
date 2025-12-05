package main

type Node struct {
	cursor int
	values []int

	next *Node
}

func (n *Node) IsFull() bool {
	return n.cursor == len(n.values)
}

func (n *Node) Push(v int) {
	n.values[n.cursor] = v
	n.cursor++
}

func (n *Node) Find(v int) *int {
	for _, value := range n.values {
		if value == v {
			return &v
		}
	}

	return nil
}

func createNode(capacity int) *Node {
	return &Node{
		cursor: 0,
		values: make([]int, capacity),
	}
}

type UnrolledLinkedList struct {
	capacity int
	head     *Node
}

func NewUnrolledLinkedList(capacity int) *UnrolledLinkedList {
	return &UnrolledLinkedList{
		capacity: capacity,
		head:     createNode(capacity),
	}
}

func (l *UnrolledLinkedList) Push(v int) {
	var prev, current *Node
	current = l.head

	for current != nil {
		if !current.IsFull() {
			break
		}

		prev = current
		current = current.next
	}

	if current == nil {
		current = createNode(l.capacity)
		prev.next = current
	}

	current.Push(v)
}

func (l *UnrolledLinkedList) Exist(v int) bool {
	current := l.head

	for current != nil {
		if current.Find(v) != nil {
			return true
		}

		current = current.next
	}

	return false
}
