package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InOrder traversal algorithm: left-root-right
// TC: O(N)
func InorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		helper(node.Left)
		result = append(result, node.Val)
		helper(node.Right)
	}

	helper(root)

	return result
}
