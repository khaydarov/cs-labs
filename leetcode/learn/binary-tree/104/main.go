package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Queue struct {
	data []*TreeNode
}

func (q *Queue) Enqueue(node *TreeNode) {
	q.data = append(q.data, node)
}

func (q *Queue) Dequeue() *TreeNode {
	if len(q.data) == 0 {
		return nil
	}

	node := q.data[0]
	q.data = q.data[1:]

	return node
}

func (q *Queue) Length() int {
	return len(q.data)
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

// Using top-down approach
//
// TC: O(V) where V is the number of nodes
// SC: O(V)
func maxDepth1(root *TreeNode) int {
	var result int
	var traversal func(node *TreeNode, depth int)

	traversal = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			result = max(result, depth)
		}

		traversal(node.Left, depth+1)
		traversal(node.Right, depth+1)
	}

	traversal(root, 1)

	return result
}

// Using BFS
//
// TC: O(V)
// SC: O(V)
func maxDepth2(root *TreeNode) int {
	q := Queue{}
	q.Enqueue(root)

	steps := 0
	for !q.Empty() {
		l := q.Length()

		for i := 0; i < l; i++ {
			node := q.Dequeue()

			if node.Left != nil {
				q.Enqueue(node.Left)
			}

			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}

		steps++
	}

	fmt.Println(steps)
	return steps
}

// Using bottom-up approach
//
// TC: O(V)
// SC: O(V)
func maxDepth3(root *TreeNode) int {
	var traversal func(node *TreeNode, depth int) int
	traversal = func(node *TreeNode, depth int) int {
		if node == nil {
			return depth - 1
		}

		left := traversal(node.Left, depth+1)
		right := traversal(node.Right, depth+1)

		return max(left, right)
	}

	result := traversal(root, 1)

	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
