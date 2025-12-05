package main

import (
	"fmt"
	"testing"
)

func TestXorLinkedList(t *testing.T) {
	l := NewXorLinkedList()
	l.AddAtTail(1)
	l.AddAtTail(2)
	l.AddAtTail(3)
	l.AddAtTail(4)

	var result []int
	l.Iterate(func(node *Node) {
		result = append(result, node.value)
	}, forwardIterationType)

	fmt.Println(result)
}
