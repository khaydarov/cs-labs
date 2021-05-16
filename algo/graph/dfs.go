package main

import "fmt"

// DFS
// TC: O(V + E), where V is the number of vertices and E is the number of edges in the graph.
func DFS(graph [][]int, x, y int) bool {
	if x == y {
		return true
	}

	for _, path := range graph {
		if path[0] == x {
			if DFS(graph, path[1], y) == true {
				return true
			}
		}
	}

	return false
}

// DFS1
// TC: O(V + E), where V is the number of vertices and E is the number of edges in the graph.
// SC: O(V)
func DFS1(graph [][]int, x, y int) bool {
	var visited [100][100]int

	var dfsUtil func(i, j int) bool
	dfsUtil = func(i, j int) bool {
		if i == j {
			return true
		}

		visited[i][j] = 1
		for _, path := range graph {
			if path[0] == i {
				if dfsUtil(path[1], j) == true {
					return true
				}
			}
		}
		return false
	}

	return dfsUtil(x, y)
}

// Implementation with Stack
// -----------
//

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

	return s.data[l - 1]
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

func DFS3(graph [][]int, x, y int) bool {
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

func main() {
	graph := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{3, 5},
		{3, 6},
		{4, 7},
		{6, 7},
	}

	r := DFS3(graph, 5, 5)
	fmt.Println(r)
}