package main

import "fmt"

// Node
// Time complexity
//  - Access: O(N) | O(N)
//  - Search: O(N) | O(N)
//  - Insertion: O(1) | O(1)
//  - Deletion: O(1) | O(1)
//
// Space complexity: O(N)
type Node struct {
	Val 	int
	Prev 	*Node
	Next 	*Node
}

type DoublyLinkedList struct {
	head 	*Node
	length 	int
}

func (list *DoublyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}

	if list.head == nil {
		list.head = node
	} else {
		node.Next = list.head
		list.head.Prev = node
		list.head = node
	}

	list.length++
}

func (list *DoublyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}
	if list.head == nil {
		list.head = node
	} else {
		current := list.head
		for current.Next != nil {
			current = current.Next
		}

		current.Next = node
		node.Prev = current
	}

	list.length++
}

func (list *DoublyLinkedList) AddAtIndex(index, val int) {
	if index < 0 || index > list.length {
		return
	}

	if index == 0 {
		list.AddAtHead(val)
	} else if index == list.length {
		list.AddAtTail(val)
	} else {
		current := list.head
		for index > 0 {
			current = current.Next
			index--
		}

		node := &Node{Val: val}

		// bind previous node with new one
		current.Prev.Next = node
		node.Prev = current.Prev

		// bind current node with new one
		node.Next = current
		current.Prev = node
	}
}

func (list *DoublyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > list.length {
		return
	}

	if index == 0 {
		next := list.head.Next
		next.Prev = nil
		list.head.Next = nil
		list.head = next
	} else if index == list.length - 1 {
		current := list.head
		for current.Next != nil {
			current = current.Next
		}

		current.Prev.Next = nil
		current.Prev = nil
	} else {
		current := list.head
		for index > 0 {
			current = current.Next
			index--
		}

		current.Prev.Next = current.Next
		current.Next.Prev = current.Prev
	}

	list.length--
}

func (list *DoublyLinkedList) Get(index int) *Node {
	if index < 0 || index > list.length {
		return nil
	}

	current := list.head
	for index > 0 {
		current = current.Next
		index--
	}

	return current
}

func (list *DoublyLinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}

func main() {
	l := &DoublyLinkedList{}
	l.AddAtTail(0)
	l.AddAtTail(1)
	l.AddAtTail(2)
	l.AddAtTail(3)

	l.DeleteAtIndex(1)
	l.Display()
}

