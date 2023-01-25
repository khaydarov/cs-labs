package main

import (
	"strconv"
)

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
		return 0
	}

	return s.data[l-1]
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

func evalRPN(tokens []string) int {
	stack := Stack{}

	for _, v := range tokens {
		if v != "+" && v != "-" && v != "/" && v != "*" {
			i, _ := strconv.Atoi(v)
			stack.Push(i)
		} else {
			operand2 := stack.Top()
			stack.Pop()
			operand1 := stack.Top()
			stack.Pop()

			//fmt.Println(operand1, operand2, v)
			var result int
			if v == "+" {
				result = operand1 + operand2
			} else if v == "*" {
				result = operand1 * operand2
			} else if v == "/" {
				result = operand1 / operand2
			} else if v == "-" {
				result = operand1 - operand2
			}

			stack.Push(result)
		}
	}

	return stack.Top()
}
