package main

type Node struct {
	Val 	int
	Left 	*Node
	Right 	*Node
	Next 	*Node
}

// TC: O(Log N * Log N)
// SC: O(Log N)
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	left := root.Left
	right := root.Right

	for left != nil {
		left.Next = right
		left = left.Right
		right = right.Left
	}

	connect(root.Left)
	connect(root.Right)

	return root
}

func main() {
	connect(nil)
}

