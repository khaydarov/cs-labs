package main

// Stack that stores data in array
// LIFO data structure
type Stack struct {
	data []int
}

// Push adds element to the end of stack storage
func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

// Top returns the last element from storage
func (s *Stack) Top() int {
	l := len(s.data)

	if l == 0 {
		return -1
	}

	return s.data[l - 1]
}

// Pop removes the last element from storage
func (s *Stack) Pop() bool {
	l := len(s.data)

	if l == 0 {
		return false
	}

	s.data = s.data[:l-1]
	return true
}

// Empty returns true if slack is empty
func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func main() {
}