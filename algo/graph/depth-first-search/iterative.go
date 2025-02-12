package main

// Implementation with Iteration
// TC: O(V + E), where V is the number of vertices and E is the number of edges in the graph.
// SC: O(V)
func Iterative(graph [][]int, x, y int) bool {
	s := NewStack()
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
