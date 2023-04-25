package main

import (
	"fmt"
	"strconv"
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

	one.Next = two
	two.Next = three
	three.Next = four

	return one
}

func testCase3() *ListNode {
	one := &ListNode{Val: 1}
	two := &ListNode{Val: 2}
	three := &ListNode{Val: 3}

	one.Next = two
	two.Next = three

	return one
}

var testCases = []struct {
	head     *ListNode
	expected string
}{
	{
		testCase1(),
		"1.3.5.2.4.",
	},
	{
		testCase2(),
		"1.3.2.4.",
	},
	{
		testCase3(),
		"1.3.2.",
	},
}

func serialize(node *ListNode) string {
	var result string
	for node != nil {
		result += strconv.Itoa(node.Val) + "."
		node = node.Next
	}

	return result
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			head := oddEvenList(testCase.head)
			answer := serialize(head)
			if serialize(head) != testCase.expected {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expected, answer)
			}
		})
	}
}
