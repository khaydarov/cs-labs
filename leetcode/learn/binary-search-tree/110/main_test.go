package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	root   *TreeNode
	expect bool
}{
	{
		testCase(),
		false,
	},
}

func testCase() *TreeNode {
	zero := &TreeNode{0, nil, nil}
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{2, nil, nil}
	three := &TreeNode{3, nil, nil}
	four := &TreeNode{4, nil, nil}
	five := &TreeNode{5, nil, nil}
	six := &TreeNode{6, nil, nil}
	seven := &TreeNode{7, nil, nil}

	zero.Left = one

	one.Left = two
	one.Right = three

	two.Left = four
	two.Right = five

	three.Left = six
	three.Right = seven

	return zero
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := isBalanced(testCase.root)
			if output != testCase.expect {
				t.Errorf("Output is not equals to expected")
			}
		})
	}
}
