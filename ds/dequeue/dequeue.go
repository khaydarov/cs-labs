package main

type Dequeue interface {
	// PushFront inserts new value in front of Dequeue: before queue
	PushFront(x int)

	// PushBack inserts new value at the end of Dequeue: last queue element
	PushBack(x int)

	// IsFull return true or false if Dequeue is full
	IsFull() bool

	// IsEmpty return true or false if Dequeue is empty
	IsEmpty() bool

	GetSize() int
	GetFront() int
	GetBack() int
	PopFront() int
	PopBack() int
}
