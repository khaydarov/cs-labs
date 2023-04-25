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
	list1  *ListNode
	list2  *ListNode
	expect string
}{
	{
		makeList([]int{1, 2, 3}),
		makeList([]int{1}),
		"1.1.2.3.",
	},
	{
		makeList([]int{1, 2, 3}),
		makeList([]int{4, 5, 6}),
		"1.2.3.4.5.6.",
	},
	{
		nil,
		nil,
		"",
	},
	{
		makeList([]int{1}),
		nil,
		"1.",
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := mergeTwoLists(testCase.list1, testCase.list2)
			sOutput := serialize(output)

			if sOutput != testCase.expect {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, sOutput)
			}
		})
		//fmt.Println(i, testCase)
	}
}
