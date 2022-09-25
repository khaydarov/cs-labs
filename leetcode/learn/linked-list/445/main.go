package _45

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

func Reverse(l *ListNode) *ListNode {
	if l.Next == nil {
		return l
	}

	rest := Reverse(l.Next)
	l.Next.Next = l
	l.Next = nil

	return rest
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1 := Length(l1)
	len2 := Length(l2)

	if len2 > len1 {
		return addTwoNumbers(l2, l1)
	}

	l3, carry := rec(l1, l2, 0, len1-len2)

	result := l3
	if carry != 0 {
		result = &ListNode{Val: carry, Next: l3}
	}

	return result
}

func rec(l1, l2 *ListNode, p1, p2 int) (*ListNode, int) {
	var sum, carry int
	if l1.Next == nil && l2.Next == nil {
		sum := l1.Val + l2.Val
		carry = sum / 10
		sum %= 10

		return &ListNode{Val: sum}, carry
	}

	var l3 *ListNode
	if p2 > p1 {
		l3, carry = rec(l1.Next, l2, p1+1, p2)
		sum = l1.Val + carry
	} else {
		l3, carry = rec(l1.Next, l2.Next, p1+1, p2+1)
		sum = l1.Val + l2.Val + carry
	}

	carry = sum / 10
	sum %= 10

	return &ListNode{Val: sum, Next: l3}, carry
}

func main() {
	l1 := &ListNode{9, nil}
	//l1.Add(9)
	//l1.Add(9)
	//l1.Add(9)
	//l1.Add(9)
	//l1.Add(9)
	//l1.Add(9)

	l2 := &ListNode{9, nil}
	//l2.Add(9)
	//l2.Add(9)
	//l2.Add(9)

	ans := addTwoNumbers(l1, l2)
	Display(ans)
}
