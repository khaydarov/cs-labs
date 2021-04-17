package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
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

func main() {
	one := &ListNode{3, nil}
	two := &ListNode{2, nil}
	three := &ListNode{2, nil}
	four := &ListNode{5, nil}
	//five := &ListNode{3, nil}

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = one

	fmt.Println(hasCycle(one))
}