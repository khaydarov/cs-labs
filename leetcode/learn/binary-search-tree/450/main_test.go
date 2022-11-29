package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	root            *TreeNode
	key             int
	expectRootValue int
}{
	{
		root:            getTree(),
		key:             4,
		expectRootValue: 5,
	},
	{
		root:            getTree(),
		key:             1,
		expectRootValue: 4,
	},
}

func getTree() *TreeNode {
	zero := &TreeNode{4, nil, nil}
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{6, nil, nil}
	three := &TreeNode{5, nil, nil}
	four := &TreeNode{8, nil, nil}

	zero.Left = one
	zero.Right = two

	two.Left = three
	two.Right = four

	return zero
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := deleteNode(testCase.root, testCase.key)

			if output.Val != testCase.expectRootValue {
				t.Errorf("output is not equals to expected")
			}
		})
	}
}
