package main

// TC: O(N)
// SC: O(N)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(root, min, max *TreeNode) *TreeNode
	dfs = func (node, min, max *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val > max.Val {
			return dfs(node.Left, min, max)
		} else if node.Val < min.Val {
			return dfs(node.Right, min, max)
		}

		return node
	}

	if p.Val > q.Val {
		return dfs(root, q, p)
	}

	return dfs(root, p, q)
}
