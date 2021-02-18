package leetcode

import "fmt"

type Node struct {
	Val 	int
	Prev 	*Node
	Next 	*Node
	Child	*Node
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

func Display(l *Node) {
	for l != nil {
		fmt.Println(l)
		l = l.Next
	}
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	current := &Node{}
	head := current
	for root != nil {
		current.Next = root
		root.Prev = current
		current = current.Next

		if root.Child != nil {
			next := root.Next

			rest := flatten(root.Child)
			current.Next = rest
			rest.Prev = current

			for current.Next != nil {
				current = current.Next
			}

			root.Child = nil
			root = next
		} else {
			root = root.Next
		}
	}

	head.Next.Prev = nil
	return head.Next
}

func makeList(arr []int) *Node {
	dummy := &Node{}
	head := dummy
	for _, v := range arr {
		dummy.Next = &Node{Val: v, Prev: dummy}
		dummy = dummy.Next
	}

	head.Next.Prev = nil
	return head.Next
}

func get(list *Node, index int) *Node {
	for index > 0 {
		list = list.Next
		index--
	}

	return list
}
func main() {

	first := makeList([]int{1, 2, 3, 4, 5, 6})
	second := makeList([]int{7, 8, 9, 10})
	third := makeList([]int{11, 12})

	get(first, 2).Child = second
	get(second, 1).Child = third

	//Display(get(first, 2))
	result := flatten(first)
	Display(result)
}

