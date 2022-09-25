package _48

import (
	"fmt"
)

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

func main() {
	l1 := &ListNode{2, nil}
	//l1.Add(1)
	//l1.Add(4)
	//l1.Add(3)
	//l1.Add(6)
	//l1.Add(5)

	sorted := MergeSort(l1)
	Display(sorted)

}
