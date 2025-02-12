package main

// PostOrder traversal algorithm: left-right-root
// TC: O(N)
// SC: O(N)
func PostOrderTravelsalRecursive(root *TreeNode) []int {
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
		helper(node.Right)
		result = append(result, node.Val)
	}

	helper(root)
	return result
}
