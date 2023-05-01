package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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
