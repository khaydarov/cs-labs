package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// TC: O(N)
// SC: O(N)
func isValidBST(root *TreeNode) bool {
	var validate func (node *TreeNode, low, high int) bool
	validate = func(node *TreeNode, low, high int) bool {
		if node == nil {
			return true
		}

		if node.Val <= low || node.Val >= high {
			return false
		}

		return validate(node.Right, node.Val, high) && validate(node.Left, low, node.Val)
	}

	high := math.MaxInt64
	low := math.MinInt64
	return validate(root, low, high)
}

type TreeNodeWithBoundary struct {
	node *TreeNode
	low, high int
}

type Stack struct {
	nodes []*TreeNodeWithBoundary
}

func (s *Stack) Push(node *TreeNode, low, high int) {
	data := &TreeNodeWithBoundary{node, low, high}
	s.nodes = append(s.nodes, data)
}

func (s *Stack) Pop() *TreeNodeWithBoundary {
	if s.Empty() {
		return nil
	}

	pop := s.nodes[len(s.nodes) - 1]
	s.nodes = s.nodes[:len(s.nodes) - 1]

	return pop
}

func (s *Stack) Empty() bool {
	return len(s.nodes) == 0
}

// TC: O(N)
// SC: O(N)
func isValidBST1(root *TreeNode) bool {
	high := math.MaxInt64
	low := math.MinInt64

	stack := &Stack{}
	stack.Push(root, low, high)

	for !stack.Empty() {
		current := stack.Pop()

		if current.node.Val <= current.low || current.node.Val >= current.high {
			return false
		}

		if current.node.Left != nil {
			stack.Push(current.node.Left, current.low, current.node.Val)
		}

		if current.node.Right != nil {
			stack.Push(current.node.Right, current.node.Val, current.high)
		}
	}

	return true
}

// InOrder recursive solution
//
// TC: O(N)
// SC: O(N)
func isValidBST2(root *TreeNode) bool {
	var prev *TreeNode
	var inorder func(node *TreeNode) bool

	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !inorder(node.Left) {
			return false
		}

		if prev != nil && prev.Val >= node.Val {
			return false
		}

		prev = node
		return inorder(node.Right)
	}

	return inorder(root)
}

func main() {
	zero := &TreeNode{5, nil, nil}
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{4, nil, nil}
	three := &TreeNode{3, nil, nil}
	four := &TreeNode{6, nil, nil}

	zero.Left = one
	zero.Right = two

	two.Left = three
	two.Right = four

	fmt.Println(isValidBST2(zero))
}