package main

import "fmt"

// NewArrayDequeue creates new instance of ArrayDequeue
func NewArrayDequeue(capacity int) ArrayDequeue {
	return ArrayDequeue{
		-1,
		0,
		capacity,
		0,
		make([]int, capacity),
	}
}

// ArrayDequeue is an implementation of Dequeue data structure using circular array
type ArrayDequeue struct {
	front    int
	rear     int
	capacity int
	length   int
	data     []int
}

func (d *ArrayDequeue) IsFull() bool {
	return (d.front == 0 && d.rear == d.capacity-1) || (d.front == d.rear+1)
}

func (d *ArrayDequeue) IsEmpty() bool {
	return d.front == -1
}

func (d *ArrayDequeue) PushFront(x int) {
	if d.IsFull() {
		fmt.Println("Overflow")
		return
	}

	if d.front == -1 {
		d.front = 0
		d.rear = 0
	} else if d.front == 0 {
		d.front = d.capacity - 1
	} else {
		d.front--
	}

	d.length++
	d.data[d.front] = x
}

func (d *ArrayDequeue) PushBack(x int) {
	if d.IsFull() {
		fmt.Println("Overflow")
		return
	}

	if d.front == -1 {
		d.front = 0
		d.rear = 0
	} else if d.rear == d.capacity-1 {
		d.rear = 0
	} else {
		d.rear++
	}

	d.length++
	d.data[d.rear] = x
}

func (d *ArrayDequeue) GetFront() int {
	if d.IsEmpty() {
		fmt.Println("Empty")
		return -1
	}

	return d.data[d.front]
}

func (d *ArrayDequeue) GetBack() int {
	if d.IsEmpty() {
		fmt.Println("Empty")
		return -1
	}

	return d.data[d.rear]
}

func (d *ArrayDequeue) PopFront() int {
	if d.IsEmpty() {
		fmt.Println("Underflow")
		return -1
	}

	popValue := d.data[d.front]
	if d.front == d.rear {
		d.front = -1
		d.rear = 0
	} else if d.front == d.capacity-1 {
		d.front = 0
	} else {
		d.front++
	}

	d.length--
	return popValue
}

func (d *ArrayDequeue) PopBack() int {
	if d.IsEmpty() {
		fmt.Println("Underflow")
		return -1
	}

	popValue := d.data[d.rear]
	if d.front == d.rear {
		d.front = -1
		d.rear = 0
	} else if d.rear == 0 {
		d.rear = d.capacity - 1
	} else {
		d.rear--
	}

	d.length--
	return popValue
}
