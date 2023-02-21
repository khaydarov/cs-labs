package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var count int
	var isUnivalSubtree func(node *TreeNode) bool

	isUnivalSubtree = func(node *TreeNode) bool {
		if node.Left == nil && node.Right == nil {
			count++
			return true
		}

		isUnival := true
		if node.Left != nil {
			isUnival = isUnivalSubtree(node.Left) && node.Left.Val == node.Val
		}

		if node.Right != nil {
			isUnival = isUnivalSubtree(node.Right) && isUnival && node.Right.Val == node.Val
		}

		if isUnival {
			count++
		}

		return isUnival
	}

	isUnivalSubtree(root)

	return count
}
