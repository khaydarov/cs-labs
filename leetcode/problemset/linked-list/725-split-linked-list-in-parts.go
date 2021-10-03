package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

// TC: O(N + k)
// SC: O(k)
func splitListToParts(head *ListNode, k int) []*ListNode {
	if head == nil {
		return make([]*ListNode, k)
	}

	n := Length(head)
	var result []*ListNode

	if k > n {
		current := head
		next := current.Next

		for current != nil {
			result = append(result, current)
			current.Next = nil
			current = next

			if next != nil {
				next = next.Next
			}
		}

		for i := 0; i < k - n; i++ {
			result = append(result, nil)
		}
	} else {
		current := head
		for k != 0 {
			m := 0
			if n % k == 0 {
				m = n / k
			} else {
				m = n / k + 1
			}

			fmt.Println(n, m, k)
			n -= m
			k--

			result = append(result, current)
			prev := current
			for m != 0 {
				prev = current
				current = current.Next
				m--
			}

			if prev != nil {
				prev.Next = nil
			}
		}
	}

	return result
}

func Length(l *ListNode) int {
	length := 0
	current := l
	for current != nil {
		length++
		current = current.Next
	}

	return length
}

func Print(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func main() {
	one := &ListNode{Val: 1}
	//two := &ListNode{Val: 2}
	//three := &ListNode{Val: 3}
	//four := &ListNode{Val: 4}

	//one.Next = two
	//two.Next = three
	//three.Next = four

	ans := splitListToParts(one, 2)
	fmt.Println(ans)
	Print(ans[0])
}