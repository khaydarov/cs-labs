package main

import "fmt"

// Linked list implementation
// Time complexity
//  - Access: O(N) | O(N)
//  - Search: O(N) | O(N)
//  - Insertion: O(1) | O(1)
//  - Deletion: O(1) | O(1)
//
// Space complexity: O(N)

type Node struct {
	value 	int
	next 	*Node
}

type LinkedList struct {
	length 	int
	head 	*Node
	tail	*Node
}

func (l *LinkedList) Len() int {
	return l.length
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (l *LinkedList) Get(index int) int {
	if index >= l.length || index < 0 {
		return -1
	}

	pointer := 0
	current := l.head
	for pointer < index {
		pointer++
		current = current.next
	}

	return current.value
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (l *LinkedList) AddAtHead(val int)  {
	node := &Node{value: val}
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
	l.length++
}


/** Append a node of value val to the last element of the linked list. */
func (l *LinkedList) AddAtTail(val int)  {
	node := &Node{value: val}
	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	l.length++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (l *LinkedList) AddAtIndex(index int, val int)  {
	if index < 0 || index > l.length {
		return
	}

	if index == 0 {
		l.AddAtHead(val)
	} else if index == l.length {
		l.AddAtTail(val)
	} else {
		pointer := 0
		current := l.head
		prev := current
		for pointer < index {
			pointer++
			prev = current
			current = current.next
		}

		node := &Node{value: val}
		node.next = current
		prev.next = node
		l.length++
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (l *LinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= l.length {
		return
	}

	if index == 0 {
		l.head = l.head.next
		l.length--
	} else {
		pointer := 0
		current := l.head
		prev := current
		for pointer < index {
			pointer++
			prev = current
			current = current.next
		}

		prev.next = current.next
		l.length--
	}
}

func (l *LinkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

