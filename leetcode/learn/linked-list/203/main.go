package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// TC: O(N)
// SC: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Val: 0}
	current := dummy

	for head != nil {
		next := head.Next
		if head.Val != val {
			current.Next = head
			head.Next = nil

			current = current.Next
		}

		head = next
	}

	return dummy.Next
}

// TC: O(N)
// SC: O(1)
func removeElements1(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Val: 0}
	dummy.Next = head
	prev, current := dummy, head

	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = current
		}

		current = current.Next
	}

	return dummy.Next
}
