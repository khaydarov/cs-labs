package geekforgeeks

import "fmt"

type ListNode struct {
	Val 	int
	Next 	*ListNode
	Down	*ListNode
}

func (l *ListNode) Add(val int) {
	if l.Next == nil {
		l.Next = &ListNode{val, nil, nil}
	} else {
		current := (*l).Next
		for current.Next != nil {
			current = current.Next
		}

		current.Next = &ListNode{val, nil, nil}
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

func constructLinkedMatrix(matrix [][]int) *ListNode {
	head := createList(matrix[0])
	current := head

	for i := 1; i < len(matrix); i++ {
		row := createList(matrix[i])

		cursor := current
		next := row
		for cursor != nil {
			cursor.Down = next

			cursor = cursor.Next
			next = next.Next
		}

		current = row
	}

	return head
}

func createList(values []int) *ListNode {
	var list *ListNode
	for i, v := range values {
		if i == 0 {
			list = &ListNode{v, nil, nil}
		} else {
			list.Add(v)
		}
	}

	return list
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	constructLinkedMatrix(matrix)
	//head := &ListNode{1, nil, nil}
	//head.Add(2)
	//head.Add(3)
	//head.Add(2)
	//head.Add(1)
	//head.Add(1)
	//head.Add(6)
	//head.Add(4)
	//head.Add(7)

	//fmt.Println(isPalindrome(head))
	//Display(head)
}