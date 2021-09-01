package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	h1 := calcHeight(root.Left, 0)
	h2 := calcHeight(root.Right, 0)

	fmt.Println(root.Left, h1)
	fmt.Println(root.Right, h2)

	if abs(h1 - h2) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}

	return x
}

func calcHeight(root *TreeNode, height int) int {
	if root == nil {
		return height
	}

	h1 := calcHeight(root.Left, height + 1)
	h2 := calcHeight(root.Right, height + 1)

	if h1 < h2 {
		return h2
	}

	return h1
}

func main() {
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

	fmt.Println(isBalanced(zero))
}