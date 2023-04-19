package main

import (
	"testing"
)

func testCase1() (*ListNode, *ListNode) {
	one := &ListNode{Val: 1}
	two := &ListNode{Val: 2}
	three := &ListNode{Val: 3}
	four := &ListNode{Val: 4}
	five := &ListNode{Val: 5}

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = five

	return one, nil
}

func testCase2() (*ListNode, *ListNode) {
	one := &ListNode{Val: 1}
	two := &ListNode{Val: 2}
	three := &ListNode{Val: 3}
	four := &ListNode{Val: 4}
	five := &ListNode{Val: 5}

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = five
	five.Next = two

	return one, two
}

func TestSolution(t *testing.T) {
	tc1, c1 := testCase1()
	o1 := detectCycle(tc1)
	if o1 != c1 {
		t.Errorf("Wrong answer tc1")
	}

	tc2, c2 := testCase2()
	o2 := detectCycle(tc2)
	if o2 != c2 {
		t.Errorf("Wrong answer tc2")
	}
}
