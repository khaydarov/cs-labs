package main

import (
	"fmt"
	"strconv"
	"testing"
)

func makeList(arr []int) *Node {
	dummy := &Node{}
	head := dummy
	for _, v := range arr {
		dummy.Next = &Node{Val: v, Prev: dummy}
		dummy = dummy.Next
	}

	head.Next.Prev = nil
	return head.Next
}

func get(list *Node, index int) *Node {
	for index > 0 {
		list = list.Next
		index--
	}

	return list
}

func testCase1() *Node {
	first := makeList([]int{1, 2, 3, 4, 5, 6})
	second := makeList([]int{7, 8, 9, 10})
	third := makeList([]int{11, 12})

	get(first, 2).Child = second
	get(second, 1).Child = third

	return first
}

func testCase2() *Node {
	first := makeList([]int{1, 2, 3})
	second := makeList([]int{4, 5})

	get(first, 1).Child = second

	return first
}

func testCase3() *Node {
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}

	one.Child = two
	two.Child = three

	return one
}

func serialize(node *Node) string {
	var result string
	for node != nil {
		result += strconv.Itoa(node.Val) + "."
		node = node.Next
	}

	return result
}

var testCases = []struct {
	root   *Node
	expect string
}{
	{
		testCase1(),
		"1.2.3.7.8.11.12.9.10.4.5.6.",
	},
	{
		testCase2(),
		"1.2.4.5.3.",
	},
	{
		testCase3(),
		"1.2.3.",
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			node := flatten1(testCase.root)

			if testCase.expect != serialize(node) {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, serialize(node))
			}
		})
	}
}
