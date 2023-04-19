package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// TC: O(N)
// SC: O(1)
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}