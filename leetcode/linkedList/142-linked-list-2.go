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

func (l *ListNode) AddTail(pos int) {
	if pos == -1 {
		return
	}

	posNode := l
	posIndex := 0
	for posIndex < pos {
		posIndex++
		posNode = posNode.Next
	}

	tail := l
	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = posNode
}

func Display(list *ListNode) {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	p1 := head
	p2 := head

	maxLength := 10000
	for maxLength != 0 {
		// Slow pointer
		if p1.Next == nil {
			return false
		}

		p1 = p1.Next

		// Fast pointer
		if p2.Next == nil {
			return false
		}

		p2 = p2.Next
		if p2.Next == nil {
			return false
		}

		p2 = p2.Next
		if p1 == p2 {
			return true
		}
		maxLength--
	}

	return false
}

func createList(values []int, pos int) *ListNode {
	var list *ListNode
	for i, v := range values {
		if i == 0 {
			list = &ListNode{v, nil}
		} else {
			list.Add(v)
		}
	}

	if list != nil {
		list.AddTail(pos)
	}

	return list
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p1 := head
	p2 := head

	maxLength := 10000
	for maxLength != 0 {
		// Slow pointer
		if p1.Next == nil {
			return nil
		}

		p1 = p1.Next

		// Fast pointer
		if p2.Next == nil {
			return nil
		}

		p2 = p2.Next
		if p2.Next == nil {
			return nil
		}

		p2 = p2.Next
		if p1 == p2 {
			break
		}
		maxLength--
	}

	if maxLength == 0 {
		return nil
	}

	p1 = head
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}

func main() {
	testCases := []struct{
		head 	[]int
		pos 	int
	}{
		{[]int{3, 2, 0, -4}, 1},
		{[]int{1, 2}, 0},
		{[]int{1}, -1},
		{[]int{}, -1},
		{[]int{1, 1, 1}, -1},
	}

	for _, testCase := range testCases {
		list := createList(testCase.head, testCase.pos)
		fmt.Println(detectCycle(list))
		//fmt.Println(hasCycle(list))
	}
}