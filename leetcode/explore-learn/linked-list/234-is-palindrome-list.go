package leetcode

import "fmt"

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

func main() {
	head := &ListNode{1, nil}
	head.Add(2)
	head.Add(3)
	head.Add(2)
	head.Add(1)
	//head.Add(1)
	//head.Add(6)
	//head.Add(4)
	//head.Add(7)

	fmt.Println(isPalindrome(head))
	//Display(head)
}