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
	head     *ListNode
	k        int
	expected string
}{
	{
		makeList([]int{1, 2, 3, 4, 5}),
		1,
		"5.2.3.4.1.",
	},
	{
		makeList([]int{100, 90}),
		1,
		"90.100.",
	},
	{
		makeList([]int{100, 90}),
		2,
		"90.100.",
	},
	{
		makeList([]int{1, 2, 3, 4, 5, 6}),
		3,
		"1.2.4.3.5.6.",
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := swapNodes(testCase.head, testCase.k)

			if serialize(output) != testCase.expected {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expected, serialize(output))
			}
		})
	}
}
