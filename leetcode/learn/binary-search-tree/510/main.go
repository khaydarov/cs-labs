package main

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

// TC: O(H)
// SC: O(1)
func inorderSuccessor(node *Node) *Node {
	if node.Right != nil {
		return minValue(node.Right)
	}

	current := node.Parent
	for current != nil {
		if current.Right == node {
			current = current.Parent
			node = node.Parent
		} else {
			return current
		}
	}

	return nil
}

func minValue(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}

	return current
}
