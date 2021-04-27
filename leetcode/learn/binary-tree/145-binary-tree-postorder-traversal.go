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

// Iterative using two Stack. My solution
//
// TC: O(N)
// SC: O(N)
func postorderTraversal(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	s1 := Stack{}
	s2 := Stack{}
	current := root

	for current != nil || !s1.Empty() {
		for current != nil {
			s1.Push(current)
			current = current.Left
		}

		node := s1.Pop()
		if node.Right == nil {
			result = append(result, node.Val)
		} else {
			if s2.Top() == node {
				result = append(result, node.Val)
				s2.Pop()
			} else {
				s1.Push(node)
				s2.Push(node)
				current = node.Right
			}
		}
	}

	return result
}

// Iterative using two Stacks. GeekforGeeks
//
// TC: O(N)
// SC: O(N)
func postorderTraversal1(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	s1 := Stack{}
	s2 := Stack{}

	s1.Push(root)
	for !s1.Empty() {
		node := s1.Pop()
		s2.Push(node)

		if node.Left != nil {
			s1.Push(node.Left)
		}

		if node.Right != nil {
			s1.Push(node.Right)
		}
	}

	for !s2.Empty() {
		node := s2.Pop()
		result = append(result, node.Val)
	}

	fmt.Println(result)
	return result
}

// Recursive
//
// TC: O(N), where N is the amount of nodes
// SC: O(N)
func postorderTraversal3(root *TreeNode) []int {
	var treeWalker func (node *TreeNode)
	var result []int

	treeWalker = func (node *TreeNode) {
		if node == nil {
			return
		}

		treeWalker(node.Left)
		treeWalker(node.Right)
		result = append(result, node.Val)
	}

	treeWalker(root)

	fmt.Println(result)

	return result
}

func main() {
	//one := &TreeNode{1, nil, nil}
	//two := &TreeNode{4, nil, nil}
	//three := &TreeNode{2, nil, nil}
	//four := &TreeNode{6, nil, nil}
	//five := &TreeNode{7, nil, nil}
	//six := &TreeNode{5, nil, nil}
	//seven := &TreeNode{3, nil, nil}
	//eight := &TreeNode{9, nil, nil}
	//
	//one.Left = two
	//one.Right =  three
	//
	//two.Left = four
	//two.Right = five
	//
	//three.Left = six
	//three.Right = seven
	//
	//four.Left = eight


	one := &TreeNode{1, nil, nil}
	two := &TreeNode{2, nil, nil}
	three := &TreeNode{3, nil, nil}
	four := &TreeNode{4, nil, nil}
	five := &TreeNode{5, nil, nil}
	six := &TreeNode{6, nil, nil}
	seven := &TreeNode{7, nil, nil}
	eight := &TreeNode{8, nil, nil}

	one.Left = two
	one.Right = three

	two.Left = four
	two.Right = five

	five.Left = six
	five.Right = seven

	three.Right = eight

	//one := &TreeNode{1, nil, nil}
	//two := &TreeNode{2, nil, nil}
	//three := &TreeNode{3, nil, nil}
	//
	//one.Right = two
	//two.Left = three

	postorderTraversal1(one)
}