package main

type Stack struct {
	data []int
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Top() int {
	if len(s.data) == 0 {
		return -1
	}

	l := len(s.data)

	return s.data[l-1]
}

func (s *Stack) Pop() bool {
	if len(s.data) == 0 {
		return false
	}

	l := len(s.data)

	s.data = s.data[:l-1]
	return true
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func NewStack() *Stack {
	return &Stack{}
}
