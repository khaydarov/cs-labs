package main

import (
	"fmt"
	"testing"
)

var testCases = []TestCase{
	{
		testCase().root,
		testCase().p,
		testCase().q,
		testCase().expect,
	},
}

type TestCase struct {
	root   *TreeNode
	p      *TreeNode
	q      *TreeNode
	expect *TreeNode
}

func testCase() TestCase {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 3}
	four := &TreeNode{Val: 4}
	five := &TreeNode{Val: 5}
	six := &TreeNode{Val: 6}

	four.Left = two
	four.Right = five

	two.Left = one
	two.Right = three

	five.Right = six

	return TestCase{
		four,
		two,
		three,
		two,
	}
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := lowestCommonAncestor(testCase.root, testCase.p, testCase.q)

			if output.Val != testCase.expect.Val {
				t.Errorf("output is not equals to expected")
			}
		})
	}
}
