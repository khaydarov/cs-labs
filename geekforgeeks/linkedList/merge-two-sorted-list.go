package geekforgeeks

func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	l3 := &ListNode{0, nil}
	head := l3

	var min int
	for l1 != nil || l2 != nil {
		if l1 == nil && l2 != nil {
			min = l2.Val
			l2 = l2.Next
		} else if l2 == nil && l1 != nil {
			min = l1.Val
			l1 = l1.Next
		} else if l1 != nil && l2 != nil && l1.Val < l2.Val {
			min = l1.Val
			l1 = l1.Next
		} else if l1 != nil && l2 != nil && l1.Val >= l2.Val {
			min = l2.Val
			l2 = l2.Next
		}

		l3 = &ListNode{Val: min, Next: l3}
	}

	return head.Next
}
