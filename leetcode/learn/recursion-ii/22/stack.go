package main

type Stack struct {
	elements []rune
}

func (s *Stack) Push(element rune) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() rune {
	if len(s.elements) == 0 {
		return 0
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]rune, 0),
	}
}
