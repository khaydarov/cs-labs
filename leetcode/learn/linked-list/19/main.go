package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// TC: O(N)
// SC: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var prev, slow, fast *ListNode
	slow, fast = head, head
	for i := 1; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		prev = slow
		slow = slow.Next
	}

	if prev == nil {
		return slow.Next
	}

	prev.Next = prev.Next.Next

	return head
}

// TC: O(N)
// SC: O(N)
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, nil}
	dummy.Next = head

	removeFromEnd(dummy, n)

	return dummy.Next
}

func removeFromEnd(node *ListNode, n int) int {
	if node == nil {
		return 0
	}

	i := removeFromEnd(node.Next, n)
	if i == n {
		node.Next = node.Next.Next
	}

	return i + 1
}

// TC: O(N)
// SC: O(1)
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, nil}
	dummy.Next = head

	slow := dummy
	fast := dummy

	for i := 1; i <= n+1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
