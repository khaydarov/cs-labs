package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	root  *TreeNode
	count int
}{
	{
		testCase1(),
		0,
	},
	{
		testCase2(),
		4,
	},
	{
		testCase3(),
		6,
	},
}

func testCase1() *TreeNode {
	return nil
}

func testCase2() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 5}

	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 5}

	root.Right.Right = &TreeNode{Val: 5}

	return root
}

func testCase3() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 5}

	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 5}

	root.Right.Right = &TreeNode{Val: 5}

	return root
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := countUnivalSubtrees(testCase.root)
			if output != testCase.count {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.count, output)
			}
		})
	}
}
