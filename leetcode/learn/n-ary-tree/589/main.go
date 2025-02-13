package main

type Node struct {
	Val      int
	Children []*Node
}

// PreOrderRecursive returns the pre-order traversal of the n-ary tree.
// Time complexity: O(n)
// Space complexity: O(n)
func PreOrderRecursive(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}

	var helper func(node *Node)
	helper = func(node *Node) {
		if node == nil {
			return
		}

		result = append(result, node.Val)
		for _, v := range node.Children {
			helper(v)
		}
	}

	helper(root)

	return result
}

type Stack struct {
	data []*Node
}

func (s *Stack) Push(x *Node) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() *Node {
	if len(s.data) == 0 {
		return nil
	}

	last := len(s.data) - 1

	v := s.data[last]
	s.data = s.data[:last]

	return v
}

func (s *Stack) Top() *Node {
	if len(s.data) == 0 {
		return nil
	}

	return s.data[len(s.data)-1]
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func PreOrderIterative(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}

	s := Stack{}
	s.Push(root)

	for !s.Empty() {
		node := s.Pop()
		result = append(result, node.Val)

		for i := len(node.Children) - 1; i >= 0; i-- {
			s.Push(node.Children[i])
		}
	}

	return result
}
