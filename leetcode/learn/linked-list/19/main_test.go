package main

import "testing"

func testCase() (*ListNode, *ListNode, int) {
	one := &ListNode{Val: 1}
	two := &ListNode{Val: 2}
	three := &ListNode{Val: 3}
	four := &ListNode{Val: 4}
	five := &ListNode{Val: 5}

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = five

	return one, two, 4
}

func TestSolution(t *testing.T) {
	head, node, n := testCase()

	result := removeNthFromEnd(head, n)

	for result != nil {
		if result == node {
			t.Errorf("Wrong answer")
		}
		result = result.Next
	}
}
