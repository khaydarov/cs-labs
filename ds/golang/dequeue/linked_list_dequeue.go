package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type LinkedListDequeue struct {
	head     *Node
	tail     *Node
	capacity int
	length   int
}

func NewLinkedListDequeue(capacity int) LinkedListDequeue {
	head := &Node{}
	tail := &Node{}

	head.Next = tail
	tail.Prev = head

	return LinkedListDequeue{
		head,
		tail,
		capacity,
		0,
	}
}

func (l *LinkedListDequeue) IsEmpty() bool {
	return l.head.Next == l.tail && l.tail.Prev == l.head
}

func (l *LinkedListDequeue) IsFull() bool {
	return l.capacity == l.length
}

func (l *LinkedListDequeue) PushFront(x int) {
	if l.IsFull() {
		fmt.Println("Overflow")

		return
	}

	node := &Node{Val: x}
	node.Next = l.head.Next
	node.Prev = l.head

	l.head.Next.Prev = node
	l.head.Next = node

	l.length++
}

func (l *LinkedListDequeue) PushBack(x int) {
	if l.IsFull() {
		fmt.Println("Overflow")

		return
	}

	node := &Node{Val: x}
	node.Prev = l.tail.Prev
	node.Next = l.tail

	l.tail.Prev.Next = node
	l.tail.Prev = node

	l.length++
}

func (l *LinkedListDequeue) GetFront() int {
	if l.IsEmpty() {
		fmt.Println("Empty")

		return -1
	}

	front := l.head.Next

	return front.Val
}

func (l *LinkedListDequeue) GetBack() int {
	if l.IsEmpty() {
		fmt.Println("Empty")

		return -1
	}

	back := l.tail.Prev

	return back.Val
}

func (l *LinkedListDequeue) PopFront() int {
	if l.IsEmpty() {
		fmt.Println("Empty")

		return -1
	}

	popValue := l.head.Next
	l.head.Next.Next.Prev = l.head
	l.head.Next = l.head.Next.Next

	l.length--
	return popValue.Val
}

func (l *LinkedListDequeue) PopBack() int {
	if l.IsEmpty() {
		fmt.Println("Empty")

		return -1
	}

	popValue := l.tail.Prev
	l.tail.Prev.Prev.Next = l.tail
	l.tail.Prev = l.tail.Prev.Prev

	l.length--
	return popValue.Val
}

func (l *LinkedListDequeue) Size() int {
	return l.length
}

func (l *LinkedListDequeue) Display() {
	current := l.head.Next
	for current != l.tail {
		fmt.Println(current.Val)
		current = current.Next
	}
}
