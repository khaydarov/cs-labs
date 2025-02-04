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

// Implementation with Iteration
// TC: O(V + E), where V is the number of vertices and E is the number of edges in the graph.
// SC: O(V)
func Iterative(graph [][]int, x, y int) bool {
	s := Stack{}
	s.Push(x)

	var visited [10]int
	for !s.Empty() {
		v := s.Top()
		s.Pop()

		if v == y {
			return true
		}

		visited[v] = 1
		for _, path := range graph {
			if path[0] == v && visited[path[1]] == 0 {
				s.Push(path[1])
			}
		}
	}

	return false
}
