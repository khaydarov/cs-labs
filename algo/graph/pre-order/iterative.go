package main

// PreOrder traversal algorithm: root-left-right
// TC: O(N)
// SC: O(N)
func PreorderTraversalIterative(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	s := Stack{}
	s.Push(root)

	for !s.Empty() {
		r := s.Pop()

		result = append(result, r.Val)

		if r.Right != nil {
			s.Push(r.Right)
		}

		if r.Left != nil {
			s.Push(r.Left)
		}
	}

	return result
}
