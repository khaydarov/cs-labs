package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	current := dummy

	var val, remaining int
	for l1 != nil || l2 != nil {
		sum := remaining
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		val = sum % 10
		remaining = sum / 10

		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	if remaining != 0 {
		current.Next = &ListNode{Val: remaining}
	}

	return dummy.Next
}
