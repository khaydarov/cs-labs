package main

import "fmt"

// Node — Linked list
type Node struct {
	Val int
	Next *Node
}

// ListQueue — Queue that stored data with Linked list
type ListQueue struct {
	data *Node
}

// Enqueue adds new value to the end of queue
func (q *ListQueue) Enqueue(x int) {
	node := &Node{x, nil}
	if q.data == nil {
		q.data = node
	} else {
		current := q.data
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

// Front returns first element value
func (q *ListQueue) Front() int {
	if q.data == nil {
		return -1
	}

	return q.data.Val
}

// Dequeue removes returns first element and removes it
func (q *ListQueue) Dequeue() int {
	if q.data == nil {
		return -1
	}

	value := q.data.Val
	q.data = q.data.Next

	return value
}

// Shows queue elements
func (q *ListQueue) Display() {
	current := q.data
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

// ----------------------------

// ArrayQueue Queue that stores data with array
type ArrayQueue struct {
	data 	[]int
	pointer int
	length 	int
}

func (q *ArrayQueue) Enqueue(x int) {
	q.data = append(q.data, x)
	q.length++
}

func (q *ArrayQueue) Dequeue() int {
	if q.length == 0 {
		return -1
	}

	value := q.data[q.pointer]
	q.data[q.pointer] = 0
	q.pointer++
	q.length--

	return value
}

func (q *ArrayQueue) Front() int {
	if q.length == 0 {
		return -1
	}

	return q.data[q.pointer]
}

func (q *ArrayQueue) Display() {
	for i := q.pointer; i < len(q.data); i++ {
		fmt.Println(q.data[i])
	}
}

func main() {
	q1 := ListQueue{}
	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)
	q1.Dequeue()

	q2 := ArrayQueue{}
	q2.Enqueue(1)
	q2.Enqueue(2)
	q2.Enqueue(3)
	q2.Dequeue()
	//q2.Display()
}