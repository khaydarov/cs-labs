package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Val: 0}

	l := calcLength(head)
	k %= l

	if k == 0 {
		return head
	}

	fast := head
	for k > 0 {
		fast = fast.Next
		k--
	}

	slow := head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	dummy.Next = slow.Next
	slow.Next = nil

	current := dummy.Next
	for current.Next != nil {
		current = current.Next
	}

	current.Next = head

	return dummy.Next
}

func calcLength(node *ListNode) int {
	var result int
	current := node
	for current != nil {
		current = current.Next
		result++
	}

	return result
}
