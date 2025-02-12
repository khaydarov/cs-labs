package main

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

	return s.data[len(s.data)-1]
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func NewStack() *Stack {
	return &Stack{}
}
