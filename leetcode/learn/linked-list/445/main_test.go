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
	l1     *ListNode
	l2     *ListNode
	expect string
}{
	{
		makeList([]int{1, 2, 3}),
		makeList([]int{1}),
		"1.2.4.",
	},
	{
		makeList([]int{7, 4, 0, 7}),
		makeList([]int{9}),
		"7.4.1.6.",
	},
	{
		makeList([]int{9, 9}),
		makeList([]int{1}),
		"1.0.0.",
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := addTwoNumbers1(testCase.l1, testCase.l2)

			if serialize(output) != testCase.expect {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, serialize(output))
			}
		})
	}
}
