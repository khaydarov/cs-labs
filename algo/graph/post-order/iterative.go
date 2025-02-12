package main

// PostOrder traversal algorithm: left-right-root
// TC: O(N)
// SC: O(N)
func PostOrderTravelsalIterative(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	s1 := NewStack()
	s2 := NewStack()

	var current *TreeNode
	current = root

	for current != nil || !s1.Empty() {
		for current != nil {
			s1.Push(current)
			current = current.Left
		}

		node := s1.Pop()
		if node.Right == nil {
			result = append(result, node.Val)
		} else {
			if s2.Top() == node {
				result = append(result, node.Val)
				s2.Pop()
			} else {
				s1.Push(node)
				s2.Push(node)
				current = node.Right
			}
		}
	}

	return result
}
