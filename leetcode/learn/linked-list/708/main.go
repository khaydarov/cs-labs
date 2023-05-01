package main

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		newNode := &Node{Val: x}
		newNode.Next = newNode

		return newNode
	}

	var minNode *Node
	var maxNode *Node

	minNode = aNode
	maxNode = aNode

	current := aNode.Next
	for current != aNode {
		if minNode.Val >= current.Val {
			minNode = current
		}

		if maxNode.Val <= current.Val {
			maxNode = current
		}

		current = current.Next
	}

	newNode := &Node{Val: x}

	if newNode.Val > maxNode.Val || newNode.Val < minNode.Val {
		newNode.Next = maxNode.Next
		maxNode.Next = newNode
	} else {
		for newNode.Val > minNode.Next.Val {
			minNode = minNode.Next
		}

		newNode.Next = minNode.Next
		minNode.Next = newNode
	}

	return aNode
}
