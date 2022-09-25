package main

import "fmt"

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

type Stack struct {
	data []*TreeNode
}

func (s *Stack) Push(node *TreeNode) {
	s.data = append(s.data, node)
}

func (s *Stack) Pop() *TreeNode {
	if len(s.data) == 0 {
		return nil
	}

	top := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]

	return top
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		}

		if root.Left != nil && root.Right == nil {
			return root.Left
		}

		if root.Right != nil && root.Left == nil {
			return root.Right
		}

		current := root.Right

		if current != nil {
			for current.Left != nil {
				current = current.Left
			}

			root.Val = current.Val
			root.Right = deleteNode(root.Right, current.Val)
		}
	}

	return root
}

func main() {
	zero := &TreeNode{4, nil, nil}
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{6, nil, nil}
	three := &TreeNode{5, nil, nil}
	four := &TreeNode{8, nil, nil}

	zero.Left = one
	zero.Right = two

	two.Left = three
	two.Right = four

	fmt.Println(deleteNode(zero, 6))
	fmt.Println(zero.Left, zero.Right)
}