package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	tree        *TreeNode
	isSymmetric bool
}{
	{
		tree:        getTree1(),
		isSymmetric: true,
	},
	{
		tree:        getTree2(),
		isSymmetric: false,
	},
	{
		tree:        getTree3(),
		isSymmetric: true,
	},
}

func getTree1() *TreeNode {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 2}

	one.Left = two
	one.Right = three

	four := &TreeNode{Val: 3}
	five := &TreeNode{Val: 3}

	two.Left = four
	three.Right = five

	return one
}

func getTree2() *TreeNode {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 2}

	one.Left = two
	one.Right = three

	four := &TreeNode{Val: 3}
	five := &TreeNode{Val: 3}

	two.Left = four
	three.Left = five

	return one
}

func getTree3() *TreeNode {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 2}

	one.Left = two
	one.Right = three

	four := &TreeNode{Val: 3}
	five := &TreeNode{Val: 3}

	two.Left = four
	three.Right = five

	six := &TreeNode{Val: 1}
	seven := &TreeNode{Val: 1}

	two.Right = six
	three.Left = seven

	return one
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			output := isSymmetric(testCase.tree)
			if output != testCase.isSymmetric {
				t.Errorf("Tree is symmetric, but got incorrect value")
			}
		})
	}
}
