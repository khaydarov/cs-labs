package main

type IterationType string

var (
	forwardIterationType  IterationType = "forward"
	backwardIterationType IterationType = "backward"
)

type Node struct {
	id    int
	both  int
	value int
}

type XorLinkedList struct {
	cursor int
	cache  map[int]*Node

	head *Node
	tail *Node
}

func NewXorLinkedList() *XorLinkedList {
	return &XorLinkedList{
		cursor: 1,
		cache:  make(map[int]*Node),
	}
}

func (l *XorLinkedList) AddAtHead(val int) {
	node := &Node{
		id:    l.cursor,
		both:  l.safeId(l.head),
		value: val,
	}

	if l.head != nil {
		l.head.both = l.head.both ^ node.id
	} else {
		l.tail = node
	}

	l.head = node
	l.cursor++
	l.cache[node.id] = node
}

func (l *XorLinkedList) AddAtTail(val int) {
	node := &Node{
		id:    l.cursor,
		both:  l.safeId(l.tail),
		value: val,
	}

	if l.tail != nil {
		l.tail.both = l.tail.both ^ node.id
	} else {
		l.head = node
	}

	l.tail = node
	l.cursor++
	l.cache[node.id] = node
}

func (l *XorLinkedList) RemoveFromHead() {
	if l.head == nil {
		return
	}

	next := l.cache[l.safeBoth(l.head)]

	if next != nil {
		next.both = next.both ^ l.head.id
	} else {
		l.tail = nil
	}

	delete(l.cache, l.head.id)
	l.head = next
}

func (l *XorLinkedList) RemoveFromTail() {
	if l.tail == nil {
		return
	}

	next := l.cache[l.safeBoth(l.tail)]

	if next != nil {
		next.both = next.both ^ l.tail.id
	} else {
		l.head = nil
	}

	delete(l.cache, l.tail.id)
	l.tail = next
}

func (l *XorLinkedList) Iterate(fn func(node *Node), iterationType IterationType) {
	var prev, current *Node

	if iterationType == forwardIterationType {
		current = l.head
	}

	if iterationType == backwardIterationType {
		current = l.tail
	}

	var nextId int

	for current != nil {
		fn(current)
		nextId = current.both ^ l.safeId(prev)

		prev = current
		current = l.cache[nextId]
	}
}

func (l *XorLinkedList) safeId(node *Node) int {
	if node == nil {
		return 0
	}

	return node.id
}

func (l *XorLinkedList) safeBoth(node *Node) int {
	if node == nil {
		return 0
	}

	return node.both
}
