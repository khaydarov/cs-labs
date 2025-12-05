package doubly_linked_list

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	l := Construct()
	l.AddAtTail(3)
	l.AddAtTail(4)
	l.AddAtTail(5)
	l.AddAtHead(2)
	l.AddAtIndex(0, 1)
	l.DeleteAtIndex(4)

	if l.Length() != 4 {
		t.Errorf("Incorrect length")
	}
}
