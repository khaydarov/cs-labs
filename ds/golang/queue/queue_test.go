package main

import "testing"

func TestListQueue(t *testing.T) {
	list := ConstructListQueue()
	list.Enqueue(1)
	list.Enqueue(2)
	list.Enqueue(3)

	if list.Front() != 1 {
		t.Errorf("Queue front is not correct")
	}

	list.Dequeue()
	if list.Front() != 2 {
		t.Errorf("Queue front is not correct after dequeue")
	}
}
