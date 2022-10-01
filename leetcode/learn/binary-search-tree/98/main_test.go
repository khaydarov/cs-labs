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
		testCase1(),
		false,
	},
	{
		testCase2(),
		false,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := isValidBST(testCase.root)
			if output != testCase.expect {
				t.Errorf("Output %v does not match expected %v", output, testCase.expect)
			}
		})
	}
}

func testCase1() *TreeNode {
	zero := &TreeNode{5, nil, nil}
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{4, nil, nil}
	three := &TreeNode{3, nil, nil}
	four := &TreeNode{6, nil, nil}

	zero.Left = one
	zero.Right = two

	two.Left = three
	two.Right = four

	return zero
}

func testCase2() *TreeNode {
	zero := &TreeNode{5, nil, nil}
	one := &TreeNode{4, nil, nil}
	two := &TreeNode{6, nil, nil}
	three := &TreeNode{3, nil, nil}
	four := &TreeNode{7, nil, nil}

	zero.Left = one
	zero.Right = two

	two.Left = three
	two.Right = four

	return zero
}
