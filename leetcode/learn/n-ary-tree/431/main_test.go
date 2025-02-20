package main

import (
	"testing"
)

func isSame(a *Node, b *Node) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val || len(a.Children) != len(b.Children) {
		return false
	}

	result := true
	for i := 0; i < len(a.Children); i++ {
		result = result && isSame(a.Children[0], b.Children[0])
	}

	return result
}

func buildNode() *Node {
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}
	four := &Node{Val: 4}
	five := &Node{Val: 5}
	six := &Node{Val: 6}

	one.Children = []*Node{two, three, four}
	two.Children = []*Node{five, six}

	return one
}

func TestRecursiveSolution(t *testing.T) {
	node := buildNode()
	codec := NewRecursiveCodec()

	treeNode := codec.encode(node)
	result := codec.decode(treeNode)

	if !isSame(node, result) {
		t.Error("encoded and decoded nodes are not equal")
	}
}

func TestIterativeSolution(t *testing.T) {
	node := buildNode()
	codec := NewIterativeCodec()

	treeNode := codec.encode(node)
	result := codec.decode(treeNode)

	if !isSame(node, result) {
		t.Error("encoded and decoded nodes are not equal")
	}
}
