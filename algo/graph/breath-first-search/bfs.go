package main

// BFS1 Breadth First Search
// TC: O(V + E), where V is the number of vertices and E is the number of edges in the graph
func BFS1(graph [][]int, root, target int) int {
	q := NewQueue(10)
	q.EnQueue(root)

	step := 0
	for !q.IsEmpty() {
		step++

		size := q.length
		for i := 0; i < size; i++ {
			current := q.Front()
			if current == target {
				return step
			}

			for _, path := range graph {
				if path[0] == current {
					q.EnQueue(path[1])
				}
			}

			q.DeQueue()
		}
	}

	return -1
}

// TC: O(V+E)
// SC: O(V)
func BFS(graph [][]int, root, target int) int {
	q := NewQueue(10)
	s := NewSet()

	step := 0
	q.EnQueue(root)

	for !q.IsEmpty() {
		step++
		l := q.length

		for i := 0; i < l; i++ {
			current := q.Front()

			if current == target {
				return step
			}

			for _, path := range graph {
				if path[0] == current {
					if !s.Has(path[1]) {
						q.EnQueue(path[1])
						s.Add(path[1])
					}
				}
			}

			q.DeQueue()
		}
	}

	return -1
}

func main() {

}
