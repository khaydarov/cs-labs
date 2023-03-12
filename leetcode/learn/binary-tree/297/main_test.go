package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	root       *TreeNode
	serialized string
}{
	{
		testCase1(),
		"1.2.#.#.3.4.#.#.5.#.#.",
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			c := Constructor()
			output := c.serialize(testCase.root)

			if output != testCase.serialized {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.serialized, output)
			}
		})
	}
}

func testCase1() *TreeNode {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 3}
	four := &TreeNode{Val: 4}
	five := &TreeNode{Val: 5}

	one.Left = two
	one.Right = three

	three.Left = four
	three.Right = five

	return one
}
