package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var largestNode *ListNode
	largestNode = head

	var current, next *ListNode
	current = head.Next

	for current != nil {
		next = current.Next

		var prev, insertionNode *ListNode
		insertionNode = head

		for insertionNode != current {
			if insertionNode.Val >= current.Val {
				break
			}

			prev = insertionNode
			insertionNode = insertionNode.Next
		}

		if largestNode.Val < insertionNode.Val {
			largestNode = insertionNode
		}

		largestNode.Next = current.Next
		if head == insertionNode {
			current.Next = head
			head = current
		} else if largestNode != current {
			prev.Next = current
			current.Next = insertionNode
		}

		current = next
	}

	return head
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	current := head

	for current != nil {
		prev := dummy

		for prev.Next != nil && prev.Next.Val <= current.Val {
			prev = prev.Next
		}

		next := current.Next
		current.Next = prev.Next
		prev.Next = current

		current = next
	}

	return dummy.Next
}
