package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Time complexity: O(N)
 * Space complexity: O(1)
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = Reverse(slow)
	fast = head
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}

		slow = slow.Next
		fast = fast.Next
	}
	return true
}

func Reverse(list *ListNode) *ListNode {
	if list.Next == nil {
		return list
	}

	rest := Reverse(list.Next)
	list.Next.Next = list
	list.Next = nil

	return rest
}
