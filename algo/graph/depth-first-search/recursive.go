package main

// Implementation with Recursion
// TC: O(V + E), where V is the number of vertices and E is the number of edges in the graph.
// SC: O(V)
func Recursive(graph [][]int, x, y int) bool {
	if x == y {
		return true
	}

	for _, path := range graph {
		if path[0] == x {
			if Recursive(graph, path[1], y) {
				return true
			}
		}
	}

	return false
}
