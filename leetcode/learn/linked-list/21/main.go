package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	current := dummy

	for list1 != nil || list2 != nil {
		if list1 == nil {
			current.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			current.Next = list1
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	return dummy.Next
}
