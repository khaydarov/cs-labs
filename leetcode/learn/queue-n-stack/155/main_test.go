package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	minStack := MinStack{}
	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(3)
	minStack.Push(4)

	if minStack.GetMin() != 1 {
		t.Errorf("incorrect min value")
	}

	if minStack.Top() != 4 {
		t.Errorf("incorrect top value")
	}

	minStack.Pop()

	if minStack.GetMin() != 1 {
		t.Errorf("incorrect min value after pop()")
	}

	if minStack.Top() != 3 {
		t.Errorf("incorrect top value after pop()")
	}
}
