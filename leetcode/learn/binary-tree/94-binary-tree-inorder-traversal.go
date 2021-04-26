package main

import "fmt"

type TreeNode struct {
	Val 		int
	Left, Right *TreeNode
}

type Stack struct {
	data []*TreeNode
}

func (s *Stack) Push(x *TreeNode) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() *TreeNode {
	if len(s.data) == 0 {
		return nil
	}

	last := len(s.data) - 1

	v := s.data[last]
	s.data = s.data[:last]

	return v
}

func (s *Stack) Top() *TreeNode {
	if len(s.data) == 0 {
		return nil
	}

	return s.data[len(s.data) - 1]
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

// Iterative using Stack
// TC: O(N)
// SC: O(N)
func inorderTraversal(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	s := Stack{}
	current := root

	for current != nil || !s.Empty() {
		for current != nil {
			s.Push(current)
			current = current.Left
		}

		node := s.Pop()
		result = append(result, node.Val)

		if node.Right != nil {
			current = node.Right
		}
	}

	return result
}

// Recursive
// TC: O(N), where N is the amount of nodes
// SC: O(N)
func inorderTraversal1(root *TreeNode) []int {
	var treeWalker func (node *TreeNode)
	var result []int

	treeWalker = func (node *TreeNode) {
		if node == nil {
			return
		}

		treeWalker(node.Left)
		result = append(result, node.Val)
		treeWalker(node.Right)
	}

	treeWalker(root)

	fmt.Println(result)

	return result
}

// Morris traversal
func inorderTraversal2(root *TreeNode) []int {
	var result []int
	current := root

	for current != nil {
		if current.Left == nil {
			result = append(result, current.Val)
			current = current.Right
		} else {
			preCurrent := current.Left
			for preCurrent.Right != nil {
				preCurrent = preCurrent.Right
			}

			preCurrent.Right = current
			temp := current
			current = current.Left
			temp.Left = nil
		}
	}

	fmt.Println(result)
	return result
}

func main() {
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{4, nil, nil}
	three := &TreeNode{2, nil, nil}
	four := &TreeNode{6, nil, nil}
	five := &TreeNode{7, nil, nil}
	six := &TreeNode{5, nil, nil}
	seven := &TreeNode{3, nil, nil}
	eight := &TreeNode{9, nil, nil}

	one.Left = two
	one.Right =  three

	two.Left = four
	two.Right = five

	three.Left = six
	three.Right = seven

	four.Left = eight

	inorderTraversal1(one)
}