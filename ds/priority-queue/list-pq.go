package priority_queue

import "fmt"

type Node struct {
	item Item
	next *Node
}

type ListPriorityQueue struct {
	head *Node
}

func Construct() ListPriorityQueue {
	return ListPriorityQueue{
		head: &Node{},
	}
}

func (l *ListPriorityQueue) Peek() Item {
	if l.head.next == nil {
		return Item{}
	}

	return l.head.next.item
}

func (l *ListPriorityQueue) Pop() Item {
	if l.head.next == nil {
		return Item{}
	}

	node := l.head.next
	l.head.next = node.next

	return node.item
}

func (l *ListPriorityQueue) Add(item Item) {
	newNode := &Node{
		item: item,
		next: nil,
	}

	if l.head.next == nil {
		l.head.next = newNode
		return
	}

	if newNode.item.priority > l.head.next.item.priority {
		newNode.next = l.head.next
		l.head.next = newNode
	} else {
		current := l.head.next
		for current.next != nil && current.next.item.priority > newNode.item.priority {
			current = current.next
		}

		newNode.next = current.next
		current.next = newNode
	}
}

func (l *ListPriorityQueue) Empty() bool {
	return l.head.next == nil
}

func (l *ListPriorityQueue) Print() {
	current := l.head.next
	for current != nil {
		fmt.Println(current.item)
		current = current.next
	}
}
