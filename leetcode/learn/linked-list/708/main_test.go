package main

import (
	"fmt"
	"strconv"
	"testing"
)

func testCase1() *Node {
	one := &Node{Val: 1}
	three := &Node{Val: 3}
	four := &Node{Val: 4}

	one.Next = three
	three.Next = four
	four.Next = one

	return three
}

func testCase2() *Node {
	one := &Node{Val: 1}
	one.Next = one

	return one
}

func testCase3() *Node {
	return nil
}

func testCase4() *Node {
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}
	four := &Node{Val: 4}

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = one

	return two
}

var testCases = []struct {
	head   *Node
	x      int
	expect string
}{
	{
		testCase1(),
		2,
		"3.4.1.2.",
	},
	{
		testCase2(),
		0,
		"1.0.",
	},
	{
		testCase3(),
		1,
		"1.",
	},
	{
		testCase4(),
		0,
		"2.3.4.0.1.",
	},
	{
		testCase4(),
		5,
		"2.3.4.5.1.",
	},
}

func serialize(node *Node) string {
	stop := node
	var result string
	for node != nil {
		result += strconv.Itoa(node.Val) + "."
		node = node.Next

		if node == stop {
			break
		}
	}

	return result
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			head := insert(testCase.head, testCase.x)
			if serialize(head) != testCase.expect {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, serialize(head))
			}
		})
	}
}
