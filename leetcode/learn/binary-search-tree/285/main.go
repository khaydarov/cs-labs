package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TC: O(H) in best case, O(N) in worst
// SC: O(N)
func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
	if p.Right != nil {
		return minValue(p.Right)
	}

	parents := make(map[*TreeNode]*TreeNode)
	var traverser func(node, parent *TreeNode)
	traverser = func(node, parent *TreeNode) {
		if node == nil {
			return
		}

		parents[node] = parent
		traverser(node.Left, node)
		traverser(node.Right, node)
	}

	traverser(root, nil)

	for parents[p] != nil {
		current := parents[p]
		if current.Val > p.Val {
			return current
		}

		p = parents[p]
	}

	return nil
}

func minValue(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}

	return current
}

// TC: O(H)
// SC: O(1)
func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	var result *TreeNode
	for root != nil {
		if root.Val > p.Val {
			result = root
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return result
}