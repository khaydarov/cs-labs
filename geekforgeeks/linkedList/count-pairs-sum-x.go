package geekforgeeks

import (
	"fmt"
)

type ListNode struct {
	Val 	int
	Next 	*ListNode
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

func countPairs(l1, l2 *ListNode, sum int) int {
	vals := make(map[int]bool)

	cursor := l1
	for cursor != nil {
		vals[cursor.Val] = true
		cursor = cursor.Next
	}

	count := 0
	cursor = l2
	for cursor != nil {
		search := sum - cursor.Val
		if _, ok := vals[search]; ok {
			count++
		}
		cursor = cursor.Next
	}

	return count
}

func main() {
	l1 := &ListNode{7, nil}
	l1.Add(5)
	l1.Add(1)
	l1.Add(3)

	l2 := &ListNode{3, nil}
	l2.Add(5)
	l2.Add(2)
	l2.Add(8)

	fmt.Println(countPairs(l1, l2, 10))
	//fmt.Println(isPalindrome(head))
	//Display(head)
}