package main

import "testing"

func testCase1() (*ListNode, *ListNode, *ListNode) {
	one := &ListNode{Val: 1}
	two := &ListNode{Val: 2}
	three := &ListNode{Val: 3}
	four := &ListNode{Val: 4}
	five := &ListNode{Val: 5}
	six := &ListNode{Val: 6}
	seven := &ListNode{Val: 7}

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = five
	five.Next = six
	six.Next = seven

	one2 := &ListNode{Val: 1}
	two2 := &ListNode{Val: 2}
	three2 := &ListNode{Val: 3}

	one2.Next = two2
	two2.Next = three2
	three2.Next = four

	return one, one2, four
}

func TestSolution(t *testing.T) {
	headA, headB, intersectionNode := testCase1()
	output := getIntersectionNode1(headA, headB)

	if output != intersectionNode {
		t.Errorf("Wrong answer")
	}
}
