package main

import (
	"fmt"
	"testing"
)

func testCase1() *ListNode {
	one := &ListNode{Val: 1}
	two := &ListNode{Val: 2}
	three := &ListNode{Val: 3}
	four := &ListNode{Val: 4}
	five := &ListNode{Val: 5}

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = five

	return one
}

func testCase2() *ListNode {
	one := &ListNode{Val: 1}
	two := &ListNode{Val: 2}
	three := &ListNode{Val: 3}
	four := &ListNode{Val: 4}
	five := &ListNode{Val: 5}

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = five
	five.Next = one

	return one
}

var testCases = []struct {
	head     *ListNode
	expected bool
}{
	{
		testCase1(),
		false,
	},
	{
		testCase2(),
		true,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := hasCycle(testCase.head)

			if output != testCase.expected {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expected, output)
			}
		})
	}
}
