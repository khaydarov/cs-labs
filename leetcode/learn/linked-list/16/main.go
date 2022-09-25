package _6

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Add(val int) {
	if l.Next == nil {
		l.Next = &ListNode{val, nil}
	} else {
		current := (*l).Next
		for current.Next != nil {
			current = current.Next
		}

		current.Next = &ListNode{val, nil}
	}
}

func (l *ListNode) AddNode(n *ListNode) {
	if l.Next == nil {
		l.Next = n
	} else {
		current := (*l).Next
		for current.Next != nil {
			current = current.Next
		}

		current.Next = n
	}
}

func GetTail(l *ListNode) *ListNode {
	head := l
	for head.Next != nil {
		head = head.Next
	}

	return head
}

func Length(l *ListNode) int {
	length := 0
	for l != nil {
		length++
		l = l.Next
	}

	return length
}

func Display(list *ListNode) {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

// Time Complexity: O(N)
// Space Complexity: O(1)
// But it passes two times
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, nil}
	dummy.Next = head

	length := 0
	cursor := head
	for cursor != nil {
		length++
		cursor = cursor.Next
	}

	length -= n
	cursor = dummy

	for length > 0 {
		length--
		cursor = cursor.Next
	}

	cursor.Next = cursor.Next.Next
	return dummy.Next
}

// Time Complexity: O(N)
// Space Complexity: O(1)
// It passes once
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

// Time Complexity: O(N)
// Space Complexity: O(1)
// It passes once
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, nil}
	dummy.Next = head
	removeFromEnd(dummy, n)

	return dummy.Next
}

func removeFromEnd(node *ListNode, n int) int {
	if node == nil {
		return 0
	}

	m := removeFromEnd(node.Next, n)
	if m == n {
		node.Next = node.Next.Next
	}
	return m + 1
}

func main() {
	head := &ListNode{1, nil}
	head.Add(2)
	head.Add(3)
	head.Add(4)
	head.Add(5)

	head = removeNthFromEnd(head, 5)
	Display(head)
}
