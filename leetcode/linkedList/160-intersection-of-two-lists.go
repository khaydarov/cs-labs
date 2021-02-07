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

// Time Complexity: O(N^2)
// Space Complexity: O(1)
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	for headA != nil {
		current := headB
		for current != nil {
			if headA == current {
				return current
			}

			current = current.Next
		}
		headA = headA.Next
	}
	return nil
}

// Time Complexity: O(N)
// Space Complexity: O(1)
// But it changes linked list structure
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	tailA := GetTail(headA)
	tailB := GetTail(headB)

	if tailA != tailB {
		return nil
	}

	l1 := Length(headA)
	l2 := Length(headB)
	headA = Reverse(headA)
	l3 := Length(headB)

	intersectionNode := headA
	intersectionPosition := (l1 + l2 - l3 + 1) / 2

	currentPosition := 0
	for currentPosition < intersectionPosition - 1 {
		currentPosition++
		intersectionNode = intersectionNode.Next
	}

	return intersectionNode
}

// Time Complexity: O(N)
// Space Complexity: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	tailA := GetTail(headA)
	tailB := GetTail(headB)

	if tailA != tailB {
		return nil
	}

	l1 := Length(headA)
	l2 := Length(headB)

	if l1 < l2 {
		return getIntersectionNode(headB, headA)
	}

	diff := l1 - l2
	for diff != 0{
		diff--
		headA = headA.Next
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

func Reverse(l *ListNode) *ListNode {
	var prev, current, next *ListNode

	current = l
	for current != nil {
		next = current.Next
		current.Next = prev

		prev = current
		current = next
	}
	l = prev

	return l
}

func main() {
	common := &ListNode{2, nil}

	headA := &ListNode{1, nil}
	headA.Add(9)
	headA.Add(1)
	headA.AddNode(common)
	headA.Add(4)

	headB := &ListNode{3, nil}
	headB.AddNode(common)

	ans := getIntersectionNode(headA, headB)
	fmt.Println(ans)
}