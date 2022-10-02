package main

import "testing"

func TestEmptyStack(t *testing.T) {
	s := &Stack{}
	s.Push(1)
	s.Pop()

	if !s.Empty() {
		t.Errorf("Stack must be empty")
	}
}

func TestStackTop(t *testing.T) {
	s := &Stack{}
	s.Push(1)
	if s.Top() != 1 {
		t.Errorf("Stack top value is not correct")
	}
}
