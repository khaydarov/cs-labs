package main

import (
	"fmt"
	"testing"
)

func TestUnrolledLinkedList(t *testing.T) {
	l := NewUnrolledLinkedList(3)
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)

	fmt.Println(l.Exist(3))
}
