package main

// InOrder traversal algorithm: left-root-right
// TC: O(N)
// SC: O(N)
func InorderTraversalRecursive(root *TreeNode) []int {
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
