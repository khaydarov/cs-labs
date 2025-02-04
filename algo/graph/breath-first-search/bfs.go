package main

import "fmt"

// Set structure
type Set struct {
	data   []int
	length int
}

// Add Adds new element to the set
// Time complexity: O(N)
// Note: We can use balanced tree to search with O(Log N)
func (s *Set) Add(x int) bool {
	if s.Has(x) {
		return false
	}

	s.data = append(s.data, x)
	s.length++

	return true
}

// Has Checks if Set contains element
func (s *Set) Has(x int) bool {
	for _, v := range s.data {
		if v == x {
			return true
		}
	}

	return false
}

type CircularQueue struct {
	data     []int
	head     int
	tail     int
	length   int
	capacity int
}

// EnQueue adds new element to the queue (to the end)
func (this *CircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.head = 0
		this.tail = 0
		this.data[this.head] = value
	} else {
		this.tail = (this.tail + 1) % this.capacity
		this.data[this.tail] = value
	}

	this.length++
	return true
}

// DeQueue removes element from the beginning
func (this *CircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = (this.head + 1) % this.capacity
	this.length--

	return true
}

// Front returns the first element
func (this *CircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.data[this.head]
}

// Rear returns the last element
func (this *CircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.data[this.tail]
}

// IsEmpty returns bool if queue is empty
func (this *CircularQueue) IsEmpty() bool {
	return this.length == 0
}

// IsFull returns bool if queue is full
func (this *CircularQueue) IsFull() bool {
	return this.length == this.capacity
}

// Constructor creates new CircularQueue with capacity of k
func Constructor(k int) CircularQueue {
	data := make([]int, k)

	return CircularQueue{
		data,
		-1,
		-1,
		0,
		k,
	}
}

// BFS1 Breadth First Search
// TC: O(V + E), where V is the number of vertices and E is the number of edges in the graph
func BFS1(graph [][]int, root, target int) int {
	q := Constructor(10)
	q.EnQueue(root)

	step := 0
	for q.IsEmpty() == false {
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
	q := Constructor(10)
	s := Set{}

	step := 0
	q.EnQueue(root)

	for q.IsEmpty() == false {
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

	r := BFS(graph, 1, 6)
	fmt.Println(r)
}
