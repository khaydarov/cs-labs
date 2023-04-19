package main

import (
	"fmt"
	"strconv"
	"testing"
)

func serialize(node *ListNode) string {
	var result string
	if node == nil {
		return result
	}

	current := node
	for current != nil {
		result += strconv.Itoa(current.Val) + "."
		current = current.Next
	}

	return result
}

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

	one.Next = two
	two.Next = three

	return one
}

func testCase3() *ListNode {
	one := &ListNode{Val: 1}
	two := &ListNode{Val: 2}

	one.Next = two

	return one
}

var testCases = []struct {
	head   *ListNode
	val    int
	expect string
}{
	{
		testCase1(),
		4,
		"1.2.3.5.",
	},
	{
		testCase2(),
		1,
		"2.3.",
	},
	{
		testCase3(),
		2,
		"1.",
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := removeElements(testCase.head, testCase.val)

			if serialize(output) != testCase.expect {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, serialize(output))
			}
		})
	}
}
