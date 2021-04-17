package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// InOrder traversal algorithm: left-root-right
// TC: O(N)
func inorderTraversal(root *TreeNode) []int {
	var dfs func(node *TreeNode)
	var result []int
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		result = append(result, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return result
}

// PreOrder traversal algorithm: root-left-right
// TC: O(N)
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func (node *TreeNode)

	dfs = func (node *TreeNode) {
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

// PostOrder traversal algorithm: left-right-root
// TC: O(N)
func postorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func (node *TreeNode)

	dfs = func (node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		dfs(node.Right)
		result = append(result, node.Val)
	}

	dfs(root)
	return result
}

func main() {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 3}
	four := &TreeNode{Val: 4}
	five := &TreeNode{Val: 5}
	six := &TreeNode{Val: 6}

	one.Left = two
	one.Right = three
	two.Left = four
	two.Right = five
	five.Right = six

	r := postorderTraversal(one)
	fmt.Println(r)
}