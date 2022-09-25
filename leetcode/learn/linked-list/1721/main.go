package _721

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Add(val int) {
	if l.Next == nil {
		l.Next = &ListNode{val, nil}
	} else {
		current := (*l).Next
		for current.Next != nil {
			current = current.Next
		}

		current.Next = &ListNode{val, nil}
	}
}

func (l *ListNode) AddNode(n *ListNode) {
	if l.Next == nil {
		l.Next = n
	} else {
		current := (*l).Next
		for current.Next != nil {
			current = current.Next
		}

		current.Next = n
	}
}

func GetTail(l *ListNode) *ListNode {
	head := l
	for head.Next != nil {
		head = head.Next
	}

	return head
}

func Length(l *ListNode) int {
	length := 0
	for l != nil {
		length++
		l = l.Next
	}

	return length
}

func Display(list *ListNode) {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head

	var prevX, prevY, currX, currY *ListNode
	currX = fast
	for i := 1; i < k; i++ {
		fast = fast.Next
		prevX = currX
		currX = fast
	}

	currY = slow
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next

		prevY = currY
		currY = slow
	}

	if prevX != nil {
		prevX.Next = currY
	} else {
		head = currY
	}

	if prevY != nil {
		prevY.Next = currX
	} else {
		head = currX
	}

	temp := currX.Next
	currX.Next = currY.Next
	currY.Next = temp

	return head
}

func main() {
	l1 := &ListNode{1, nil}
	l1.Add(2)
	l1.Add(3)
	l1.Add(4)
	l1.Add(5)
	l1.Add(6)

	l1 = swapNodes(l1, 1)
	Display(l1)
}
