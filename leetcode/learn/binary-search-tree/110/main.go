package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
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

	if abs(h1-h2) > 1 {
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

	h1 := calcHeight(root.Left, height+1)
	h2 := calcHeight(root.Right, height+1)

	if h1 < h2 {
		return h2
	}

	return h1
}
