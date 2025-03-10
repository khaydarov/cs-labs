package queue

import "fmt"

// Node — Doubly Linked list
type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

// ListQueue — Queue that stored data with Linked list
type ListQueue struct {
	head *Node
	tail *Node
}

func ConstructListQueue() *ListQueue {
	head := &Node{Val: 0}
	tail := &Node{Val: 0}

	head.Next = tail
	tail.Prev = head

	return &ListQueue{
		head,
		tail,
	}
}

// Enqueue adds new value to the end of queue
func (q *ListQueue) Enqueue(x int) {
	newNode := &Node{Val: x}

	lastNode := q.tail.Prev

	newNode.Prev = lastNode
	lastNode.Next = newNode

	newNode.Next = q.tail
	q.tail.Prev = newNode
}

// Front returns first element value
func (q *ListQueue) Front() int {
	if q.head.Next == q.tail {
		return -1
	}

	return q.head.Next.Val
}

// Dequeue removes returns first element and removes it
func (q *ListQueue) Dequeue() int {
	if q.head.Next == q.tail {
		return -1
	}

	value := q.head.Next.Val
	q.head.Next = q.head.Next.Next

	return value
}

// Display Shows queue elements
func (q *ListQueue) Display() {
	current := q.head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

// ----------------------------

// ArrayQueue Queue that stores data with array
type ArrayQueue struct {
	data    []int
	pointer int
	length  int
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
}
