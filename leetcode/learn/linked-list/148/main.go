package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeSort(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}

	slow := l
	fast := l

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = l
	for fast.Next != slow {
		fast = fast.Next
	}

	fast.Next = nil

	left := MergeSort(slow)
	right := MergeSort(l)

	return Merge(left, right)
}

func Merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	head := dummy
	for true {
		if l1 == nil && l2 == nil {
			break
		}

		if l1 == nil && l2 != nil {
			head.Next = l2
			l2 = l2.Next
		} else if l2 == nil && l1 != nil {
			head.Next = l1
			l1 = l1.Next
		} else if l1 != nil && l2 != nil && l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else if l2 != nil {
			head.Next = l2
			l2 = l2.Next
		}

		head = head.Next
	}

	return dummy.Next
}
