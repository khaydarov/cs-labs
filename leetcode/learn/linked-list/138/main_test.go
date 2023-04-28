package main

import (
	"fmt"
	"strconv"
	"testing"
)

func testCase() *Node {
	one := &Node{Val: 7}
	two := &Node{Val: 13}
	three := &Node{Val: 11}
	four := &Node{Val: 10}
	five := &Node{Val: 1}

	one.Next = two

	two.Next = three
	two.Random = one

	three.Next = four
	three.Random = five

	four.Next = five
	four.Random = three

	five.Random = one

	return one
}

func serialize(node *Node) string {
	var result string
	for node != nil {
		var nextStr string
		if node.Next != nil {
			nextStr = strconv.Itoa(node.Next.Val)
		} else {
			nextStr = "#"
		}

		var randomStr string
		if node.Random != nil {
			randomStr = strconv.Itoa(node.Random.Val)
		} else {
			randomStr = "#"
		}

		result += fmt.Sprintf(
			"%s->%s--%v.",
			strconv.Itoa(node.Val),
			nextStr,
			randomStr,
		)

		node = node.Next
	}

	return result
}

func TestSolution(t *testing.T) {
	node := testCase()
	expect := "7->13--#.13->11--7.11->10--1.10->1--11.1->#--7."

	copyNode := copyRandomList(node)
	serialized := serialize(copyNode)

	if serialized != expect {
		t.Errorf("Wrong answer. Expect %v but got %v", expect, serialized)
	}
}
