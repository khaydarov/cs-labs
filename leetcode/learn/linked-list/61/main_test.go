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
	k      int
	expect string
}{
	{
		makeList([]int{1, 2, 3, 4, 5}),
		2,
		"4.5.1.2.3.",
	},
	{
		makeList([]int{}),
		0,
		"",
	},
	{
		makeList([]int{1, 2, 3}),
		3,
		"1.2.3.",
	},
	{
		makeList([]int{1, 2, 3}),
		6,
		"1.2.3.",
	},
	{
		makeList([]int{1, 2, 3}),
		1,
		"3.1.2.",
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			rotated := rotateRight(testCase.head, testCase.k)
			output := serialize(rotated)

			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, output)
			}

		})
	}
}
