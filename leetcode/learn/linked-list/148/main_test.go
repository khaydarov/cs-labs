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
	list   *ListNode
	expect string
}{
	{
		makeList([]int{1, 4, 3, 2, 5}),
		"1.2.3.4.5.",
	},
	{
		makeList([]int{3, 2, 1}),
		"1.2.3.",
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := MergeSort(testCase.list)

			if serialize(output) != testCase.expect {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, serialize(output))
			}
		})
	}
}
