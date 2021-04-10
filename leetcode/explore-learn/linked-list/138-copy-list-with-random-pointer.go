package leetcode

import "fmt"

type Node struct {
	Val 	int
	Prev 	*Node
	Next 	*Node
	Random	*Node
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

// TC: O(N)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	current := head
	for current != nil {
		next := current.Next

		clone := &Node{Val: current.Val}
		current.Next = clone
		clone.Next = next

		current = next
	}

	current = head
	for current != nil {
		clone := current.Next

		if current.Random != nil {
			clone.Random = current.Random.Next
		}

		current = clone.Next
	}

	current = head
	clone := current.Next

	dummy := &Node{}
	dummy.Next = clone
	for current != nil {
		current.Next = clone.Next

		if current.Next != nil {
			clone.Next = current.Next.Next
		} else {
			clone.Next = nil
		}

		current = current.Next
		clone = clone.Next
	}

	return dummy.Next
}

// TC: O(N)
// SC: O(N)
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}

	m := make(map[*Node]*Node, 1000)

	current := head
	for current != nil {
		m[current] = &Node{Val: current.Val}
		current = current.Next
	}

	current = head
	copyHead := m[current]
	for current != nil {
		if _, ok := m[current]; !ok {
			continue
		}

		copyNode := m[current]
		copyNode.Next = m[current.Next]
		copyNode.Random = m[current.Random]

		current = current.Next
	}

	return copyHead
}

func main() {
	one := &Node{Val: 7}
	two := &Node{Val: 13}
	three := &Node{Val: 11}
	four := &Node{Val: 10}
	five := &Node{Val: 1}

	one.Next = two

	two.Next = three
	two.Random = one

	three.Next = four
	three.Random = five

	four.Next = five
	four.Random = three

	five.Random = one

	l := copyRandomList(five)
	Display(l)
	//Display(one)
}

