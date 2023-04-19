package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive approach
// TC: O(N)
// SC: O(N)
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversedHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return reversedHead
}

// Iterative approach
// TC: O(N)
// SC: O(1)
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var previous, current, next *ListNode
	current = head
	next = current.Next

	for next != nil {
		current.Next = previous
		previous = current
		current = next
		next = next.Next
	}

	current.Next = previous

	return current
}
