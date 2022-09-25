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

// Iterative approach
//
// TC: O(N)
// SC: O(N)
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if (root.Left == nil && root.Right != nil) || (root.Right == nil && root.Left != nil) {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	s1 := Stack{}
	s2 := Stack{}

	s1.Push(root.Left)
	s2.Push(root.Right)

	for !s1.Empty() && !s2.Empty() {
		left := s1.Pop()
		right := s2.Pop()

		if left.Val != right.Val {
			return false
		}

		if left.Left == nil && right.Right != nil || left.Left != nil && right.Right == nil {
			return false
		}

		if left.Right == nil && right.Left != nil || left.Right != nil && right.Left == nil {
			return false
		}

		if left.Left != nil && right.Right != nil {
			s1.Push(left.Left)
			s2.Push(right.Right)
		}

		if left.Right != nil && right.Left != nil {
			s1.Push(left.Right)
			s2.Push(right.Left)
		}
	}

	return true
}

// Recursive
//
// TC: O(N)
// SC: O(N)
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return symmetric(root.Left, root.Right)
}

func symmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
}

func main() {
	//one := &TreeNode{1, nil, nil}
	//two := &TreeNode{2, nil, nil}
	//three := &TreeNode{2, nil, nil}
	//four := &TreeNode{3, nil, nil}
	//five := &TreeNode{3, nil, nil}
	////six := &TreeNode{4, nil, nil}
	//
	//one.Left = two
	//one.Right = three
	//
	//two.Right = four
	//three.Left = five

	//five.Left = six

	one := &TreeNode{1, nil, nil}
	two := &TreeNode{2, nil, nil}
	three := &TreeNode{2, nil, nil}
	four := &TreeNode{3, nil, nil}
	five := &TreeNode{4, nil, nil}
	six := &TreeNode{4, nil, nil}
	seven := &TreeNode{3, nil, nil}

	one.Left = two
	one.Right =  three

	two.Left = four
	two.Right = five

	three.Left = six
	three.Right = seven

	fmt.Println(isSymmetric(one))
}