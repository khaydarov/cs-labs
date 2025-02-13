package main

type Node struct {
	Val      int
	Children []*Node
}

func PostOrderRecursive(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}

	var helper func(node *Node)
	helper = func(node *Node) {
		if node == nil {
			return
		}

		for _, v := range node.Children {
			helper(v)
		}

		result = append(result, node.Val)
	}

	helper(root)
	return result
}

// ----------------------------
type NodeWrapper struct {
	node   *Node
	cursor int
}

type Stack struct {
	data []*NodeWrapper
}

func (s *Stack) Push(x *NodeWrapper) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() *NodeWrapper {
	if len(s.data) == 0 {
		return nil
	}

	last := len(s.data) - 1

	v := s.data[last]
	s.data = s.data[:last]

	return v
}

func (s *Stack) Top() *NodeWrapper {
	if len(s.data) == 0 {
		return nil
	}

	return s.data[len(s.data)-1]
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func NewStack() *Stack {
	return &Stack{}
}

func PostOrderIterative(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}

	s1 := NewStack()

	var current *Node
	current = root
	for current != nil || !s1.Empty() {
		for current != nil {
			s1.Push(&NodeWrapper{current, 0})
			if len(current.Children) > 0 {
				current = current.Children[0]
			} else {
				current = nil
				break
			}
		}

		target := s1.Pop()
		if len(target.node.Children) == 0 || target.cursor == len(target.node.Children)-1 {
			result = append(result, target.node.Val)
		} else {
			target.cursor = target.cursor + 1
			s1.Push(target)
			current = target.node.Children[target.cursor]
		}
	}

	return result
}
