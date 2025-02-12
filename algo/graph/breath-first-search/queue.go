package main

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
func NewQueue(k int) CircularQueue {
	data := make([]int, k)

	return CircularQueue{
		data,
		-1,
		-1,
		0,
		k,
	}
}
