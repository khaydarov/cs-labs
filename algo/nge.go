package main

import "fmt"

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

// Optimized version of NGE
// TC: O(N)
// SC: O(1)
func nextGreaterElement(nums []int) []int {
	stack := Stack{}

	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for !stack.Empty() && stack.Top() <= nums[i] {
			stack.Pop()
		}

		if stack.Empty() {
			result[i] = -1
		} else {
			result[i] = stack.Top()
		}

		stack.Push(nums[i])
	}

	return result
}

func main() {
	nums := []int{4, 0, 1, 2, 5}

	r := nextGreaterElement(nums)
	fmt.Println(r)
}