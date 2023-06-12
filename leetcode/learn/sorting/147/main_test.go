package main

import (
	"fmt"
	"strconv"
	"testing"
)

func makeList(arr []int) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for _, v := range arr {
		dummy.Next = &ListNode{Val: v}
		dummy = dummy.Next
	}

	return head.Next
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
		makeList([]int{3, 2, 1, 5}),
		"1.2.3.5.",
	},
	{
		makeList([]int{5, 4, 3, 2, 1}),
		"1.2.3.4.5.",
	},
	{
		makeList([]int{1, 2, 2, 2}),
		"1.2.2.2.",
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := insertionSortList(testCase.head)
			serialized := serialize(output)

			if serialized != testCase.expect {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, serialized)
			}
		})
	}
}
