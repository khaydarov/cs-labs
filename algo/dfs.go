package main

import "fmt"

// TC: O(V + E), where V is the number of vertices and E is the number of edges in the graph.
func DFS1(graph [][]int, x, y int) bool {
	if x == y {
		return true
	}

	for _, path := range graph {
		if path[0] == x {
			if DFS1(graph, path[1], y) == true {
				return true
			}
		}
	}

	return false
}

// TC: O(V + E), where V is the number of vertices and E is the number of edges in the graph.
// SC: O(V)
func DFS(graph [][]int, x, y int) bool {
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

	r := DFS(graph, 4, 5)
	fmt.Println(r)
}