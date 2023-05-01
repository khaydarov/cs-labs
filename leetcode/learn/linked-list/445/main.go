package main

import (
	"../../../../ds/golang/stack"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Length(node *ListNode) int {
	var length int
	for node != nil {
		node = node.Next
		length++
	}

	return length
}

// TC: O(N1 + N2)
// SC: O(N1 + N2)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := Length(l1)
	n2 := Length(l2)

	if n2 > n1 {
		return addTwoNumbers(l2, l1)
	}

	var helper func(l1, l2 *ListNode, diff int) (*ListNode, int)
	helper = func(l1, l2 *ListNode, diff int) (*ListNode, int) {
		var sum int
		if l1.Next == nil && l2.Next == nil {
			sum = l1.Val + l2.Val

			return &ListNode{Val: sum % 10}, sum / 10
		}

		var carry int
		var node *ListNode
		if diff > 0 {
			node, carry = helper(l1.Next, l2, diff-1)
			sum = l1.Val + carry
		} else {
			node, carry = helper(l1.Next, l2.Next, diff)
			sum = l1.Val + l2.Val + carry
		}

		return &ListNode{Val: sum % 10, Next: node}, sum / 10
	}

	head, carry := helper(l1, l2, n1-n2)
	if carry != 0 {
		return &ListNode{Val: carry, Next: head}
	}

	return head
}

// --------

// Approach 2:
// Using two stacks
//
// TC: O(N1 + N2)
// SC: O(N1 + N2)
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := stack.Stack{}
	for l1 != nil {
		s1.Push(l1.Val)
		l1 = l1.Next
	}

	s2 := stack.Stack{}
	for l2 != nil {
		s2.Push(l2.Val)
		l2 = l2.Next
	}

	var head, previous *ListNode
	var carry int
	for !s1.Empty() || !s2.Empty() {
		var a, b int
		if !s1.Empty() {
			a = s1.Top()
			s1.Pop()
		}

		if !s2.Empty() {
			b = s2.Top()
			s2.Pop()
		}

		sum := a + b + carry
		carry = sum / 10

		head = &ListNode{Val: sum % 10}
		head.Next = previous

		previous = head
	}

	if carry != 0 {
		return &ListNode{Val: carry, Next: head}
	}

	return head
}
