package main

import "fmt"

type Reading struct {
	temperature, day int
}

// Stack that stores data in array
// LIFO data structure
type Stack struct {
	data []Reading
}

// Push adds element to the end of stack storage
func (s *Stack) Push(x Reading) {
	s.data = append(s.data, x)
}

// Top returns the last element from storage
func (s *Stack) Top() Reading {
	l := len(s.data)

	if l == 0 {
		return Reading{}
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
func dailyTemperatures1(T []int) []int {
	stack := Stack{}

	result := make([]int, len(T))
	for i := len(T) - 1; i >= 0; i-- {
		for !stack.Empty() && T[i] >= stack.Top().temperature {
			stack.Pop()
		}

		if stack.Empty() {
			result[i] = 0
		} else {
			result[i] = stack.Top().day - i
		}

		stack.Push(Reading{T[i], i})
	}

	return result
}

// TC: O(N)
// SC: O(N) but using also hash map
func dailyTemperatures(T []int) []int {
	stack := Stack{}
	m := make(map[int]int, len(T))

	for i, v := range T {
		for !stack.Empty() && stack.Top().temperature < v {
			m[stack.Top().day] = i - stack.Top().day
			stack.Pop()
		}
		stack.Push(Reading{v, i})
	}

	var result []int
	for i := 0; i < len(T); i++ {
		if v, ok := m[i]; ok {
			result = append(result, v)
		} else {
			result = append(result, 0)
		}
	}
	return result
}

func main() {
	nums := []int{89,62,70,58,47,47,46,76,100,70}

	r := dailyTemperatures(nums)
	fmt.Println(r)
}