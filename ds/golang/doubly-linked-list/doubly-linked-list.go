package doubly_linked_list

import "fmt"

// Node
// Time complexity
//   - Access: O(N) | O(N)
//   - Search: O(N) | O(N)
//   - Insertion: O(1) | O(1)
//   - Deletion: O(1) | O(1)
//
// Space complexity: O(N)
type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func Construct() *DoublyLinkedList {
	head := &Node{Val: 0}
	tail := &Node{Val: 0}

	head.Next = tail
	tail.Prev = head

	return &DoublyLinkedList{
		head,
		tail,
		0,
	}
}

func (list *DoublyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}
	node.Next = list.head.Next
	node.Prev = list.head

	list.head.Next.Prev = node
	list.head.Next = node
	list.length++
}

func (list *DoublyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}
	node.Next = list.tail
	node.Prev = list.tail.Prev

	list.tail.Prev.Next = node
	list.tail.Prev = node
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
		current := list.head.Next
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
	if index < 0 || index > list.length || list.length == 0 {
		return
	}

	if index == 0 {
		current := list.head.Next
		next := current.Next

		list.head.Next = next
		next.Prev = list.head
	} else if index == list.length-1 {
		current := list.tail.Prev
		previous := current.Prev

		list.tail.Prev = previous
		previous.Next = current.Next
	} else {
		current := list.head.Next
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

	current := list.head.Next
	for index > 0 {
		current = current.Next
		index--
	}

	return current
}

func (list *DoublyLinkedList) Length() int {
	return list.length
}

func (list *DoublyLinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Println(current, current.Val)
		current = current.Next
	}
}
