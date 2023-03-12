package main

import "testing"

func testCase1() (*TreeNode, *TreeNode, *TreeNode, *TreeNode) {
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

	return zero, two, three, one
}

func TestSolution(t *testing.T) {
	root, p, q, lca := testCase1()
	found := lowestCommonAncestor(root, p, q)

	if lca != found {
		t.Errorf("Wrong answer")
	}
}
