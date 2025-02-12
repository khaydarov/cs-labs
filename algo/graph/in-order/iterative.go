package main

// InOrder traversal algorithm: left-root-right
// TC: O(N)
// SC: O(N)
func InorderTraversalIterative(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	s := NewStack()
	var current *TreeNode
	current = root

	for current != nil || !s.Empty() {
		for current != nil {
			s.Push(current)
			current = current.Left
		}

		current = s.Pop()
		result = append(result, current.Val)
		current = current.Right
	}

	return result
}
