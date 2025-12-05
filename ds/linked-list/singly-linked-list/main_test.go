package singly_linked_list

import "testing"

func TestLinkedList(t *testing.T) {
	l := &LinkedList{}
	l.AddAtHead(1)
	l.AddAtTail(2)
	l.AddAtHead(0)

	if l.Get(0) != 0 {
		t.Errorf("Insertion error")
	}
}
