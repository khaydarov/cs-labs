package main

import (
	"testing"
)

func TestArrayDequeue(t *testing.T) {
	dequeue := NewArrayDequeue(10)
	dequeue.PushFront(2)
	dequeue.PushFront(1)
	dequeue.PushBack(3)
	dequeue.PushBack(4)

	if dequeue.GetFront() != 1 {
		t.Errorf("Expected front is 1")
	}

	if dequeue.GetBack() != 4 {
		t.Errorf("Expected back is 4")
	}

	dequeue.PopFront()
	if dequeue.GetFront() != 2 {
		t.Errorf("Expected front after pop is 2")
	}

	dequeue.PopBack()
	if dequeue.GetBack() != 3 {
		t.Errorf("Expected back after pop is 3")
	}
}

func TestArrayDequeueOverflow(t *testing.T) {
	dequeue := NewArrayDequeue(4)
	dequeue.PushBack(1)
	dequeue.PushBack(2)
	dequeue.PushBack(3)
	dequeue.PushBack(4)
	dequeue.PushBack(5)

	if dequeue.GetBack() != 4 {
		t.Errorf("Expected back is 4")
	}

	if dequeue.GetFront() != 1 {
		t.Errorf("Expected front is 1")
	}
}

func TestLinkedListDequeue(t *testing.T) {
	dequque := NewLinkedListDequeue(5)
	dequque.PushBack(1)
	dequque.PushBack(2)
	dequque.PushBack(3)
	dequque.PushFront(0)

	dequque.PopFront()
	dequque.PopBack()

	dequque.PushFront(11)
	dequque.PushBack(12)
	dequque.Display()

	if dequque.length != 4 {
		t.Errorf("Incorrect size value")
	}
}
