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

// TC: O(N)
// SC: O(N)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := Stack{}
	m := make(map[int]int, len(nums2))

	for _, v := range nums2 {
		for !stack.Empty() && stack.Top() < v {
			m[stack.Top()] = v
			stack.Pop()
		}
		stack.Push(v)
	}

	var result []int
	for _, v := range nums1 {
		if _, ok := m[v]; ok {
			result = append(result, m[v])
		} else {
			result = append(result, -1)
		}
	}

	return result
}

func main() {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}

	r := nextGreaterElement(nums1, nums2)
	fmt.Println(r)
}