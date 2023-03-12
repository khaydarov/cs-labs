package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Stack struct {
	data []int
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}

	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return top
}

func (s *Stack) Top() int {
	if len(s.data) == 0 {
		return -1
	}

	return s.data[len(s.data)-1]
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	s := Stack{}
	for _, v := range postorder {
		s.Push(v)
	}

	var helper func(arr []int) *TreeNode
	helper = func(arr []int) *TreeNode {
		if len(arr) == 0 || s.Empty() {
			return nil
		}

		top := s.Top()
		found := findInArray(arr, top)

		if found == -1 {
			return nil
		}

		s.Pop()

		node := &TreeNode{Val: top}
		node.Right = helper(arr[found+1:])
		node.Left = helper(arr[found+1:])

		return node
	}

	return helper(preorder)
}

func findInArray(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}

	return -1
}
