package main

import (
	"fmt"
	"strconv"
	"testing"
)

func makeList(nums []int) *ListNode {
	dummy := &ListNode{Val: 0}
	current := dummy

	for _, v := range nums {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}

	return dummy.Next
}

func serialize(node *ListNode) string {
	var result string
	for node != nil {
		result += strconv.Itoa(node.Val) + "."
		node = node.Next
	}

	return result
}

var testCases = []struct {
	head   *ListNode
	expect string
}{
	{
		makeList([]int{1, 2, 3, 4}),
		"2.1.4.3.",
	},
	{
		makeList([]int{1, 2, 3}),
		"2.1.3.",
	},
	{
		makeList([]int{1, 2}),
		"2.1.",
	},
	{
		makeList([]int{1, 2, 3, 4, 5}),
		"2.1.4.3.5.",
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			newHead := swapPairs(testCase.head)

			if serialize(newHead) != testCase.expect {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, serialize(newHead))
			}
		})
	}
}
