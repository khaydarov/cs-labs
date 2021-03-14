package main

import "fmt"

type Node struct {
	Val int
	Neighbors []*Node
}

func cloneGraph1(node *Node) *Node {
	storage := make(map[int]*Node)
	var clone func(target *Node, visited map[int]*Node) *Node

	clone = func(target *Node, visited map[int]*Node) *Node {
		if target == nil {
			return nil
		}

		if visited[target.Val] != nil {
			return visited[target.Val]
		}

		targetClone := &Node{Val: target.Val}
		visited[target.Val] = targetClone
		for _, neighbour := range target.Neighbors {
			targetClone.Neighbors = append(
				targetClone.Neighbors,
				clone(neighbour, visited),
			)
		}

		return targetClone
	}

	return clone(node, storage)
}

// Implementation with BFS
// ---------
//

type Queue struct {
	data []*Node
}

func (q *Queue) Front() *Node {
	if len(q.data) == 0 {
		return nil
	}

	return q.data[0]
}

func (q *Queue) Dequeue() {
	if len(q.data) == 0 {
		return
	}

	q.data = q.data[1:]
}

func (q *Queue) Enqueue(n *Node) {
	q.data = append(q.data, n)
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	q := Queue{}
	storage := make(map[int]*Node)
	storage[node.Val] = &Node{Val: node.Val}

	q.Enqueue(node)
	for !q.Empty() {
		l := q.Size()

		for i := 0; i < l; i++ {
			current := q.Front()

			for _, neighbour := range current.Neighbors {
				if storage[neighbour.Val] == nil {
					storage[neighbour.Val] = &Node{Val: neighbour.Val}
					q.Enqueue(neighbour)
				}
				storage[current.Val].Neighbors = append(storage[current.Val].Neighbors, storage[neighbour.Val])
			}
			q.Dequeue()
		}
	}

	return storage[node.Val]
}

func PrintGraph(node *Node) {
	if node == nil {
		return
	}
	var visited [100]int

	var dfs func(node *Node)
	dfs = func(node *Node) {
		visited[node.Val] = 1
		fmt.Println(node.Val)

		for _, v := range node.Neighbors {
			if visited[v.Val] == 0 {
				dfs(v)
			}
		}
	}

	dfs(node)
}

func main() {
	one := Node{1, nil}
	two := Node{2, nil}
	three := Node{3, nil}
	four := Node{4, nil}

	one.Neighbors = []*Node{&two, &four}
	two.Neighbors = []*Node{&one, &three}
	three.Neighbors = []*Node{&two, &four}
	four.Neighbors = []*Node{&one, &three}

	r := cloneGraph(&one)
	//fmt.Println(r)
	//PrintGraph(&one)
	PrintGraph(r)
}