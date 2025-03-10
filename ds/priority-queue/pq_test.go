package priority_queue

import (
	"testing"
)

func TestArrayPriorityQueue(t *testing.T) {
	q := ConstructArrayPriorityQueue(5)
	q.Insert(2, 2)
	q.Insert(1, 1)
	q.Insert(9, 11)
}

func TestListPriorityQueue(t *testing.T) {
	q := Construct()
	q.Add(Item{3, 3})
	q.Add(Item{1, 1})
	q.Add(Item{2, 2})

	var peek Item
	peek = q.Peek()
	if peek.value != 3 {
		t.Error("incorrect peek found", peek)
	}

	q.Add(Item{10, 10})
	peek = q.Peek()
	if peek.value != 10 {
		t.Error("incorrect peek found", peek)
	}
}
