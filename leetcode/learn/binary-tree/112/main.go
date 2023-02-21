package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return calcSum(root, root.Val, targetSum)
}

func calcSum(node *TreeNode, actualSum, targetSum int) bool {
	if isLeaf(node) {
		return actualSum == targetSum
	}

	left := false
	right := false

	if node.Left != nil {
		left = calcSum(node.Left, actualSum+node.Left.Val, targetSum)
	}

	if node.Right != nil {
		right = calcSum(node.Right, actualSum+node.Right.Val, targetSum)
	}

	return left || right
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
