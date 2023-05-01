package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return helper(head)
}

// TC: O(N)
// SC: O(N)
func helper(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}

	next := head.Next
	afterNext := next.Next

	next.Next = head
	head.Next = helper(afterNext)

	return next
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 1 - 2 - 3 - 4 -> 2 - 1 - 4 - 3
	// 1 - 2 - 3 -> 2 - 1 - 3
	var previous, current, next *ListNode
	current = head
	next = current.Next

	for current != nil && current.Next != nil {
		next = current.Next

		if previous == nil {
			head = next
		} else {
			previous.Next = next
		}

		current.Next = next.Next
		next.Next = current

		previous = current
		current = current.Next
	}

	return head
}
