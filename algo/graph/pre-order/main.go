package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrder traversal algorithm: root-left-right
// TC: O(N)
func PreorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		result = append(result, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return result
}
