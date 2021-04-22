package main

type TreeNode struct {
	Val 		int
	Left, Right *TreeNode
}

type Stack struct {
	data []*TreeNode
}

func (s *Stack) Push(x *TreeNode) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() *TreeNode {
	if len(s.data) == 0 {
		return nil
	}

	last := len(s.data) - 1

	v := s.data[last]
	s.data = s.data[:last]

	return v
}

func (s *Stack) Top() *TreeNode {
	if len(s.data) == 0 {
		return nil
	}

	return s.data[len(s.data) - 1]
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

// Iterative using Stack
// TC: O(N)
// SC: O(N)
func preorderTraversal(root *TreeNode) []int {
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

// Recursive
// TC: O(N), where N is the amount of nodes
// SC: O(N)
func preorderTraversal1(root *TreeNode) []int {
	var treeWalker func (node *TreeNode)
	var result []int

	treeWalker = func (node *TreeNode) {
		if node == nil {
			return
		}

		result = append(result, node.Val)
		treeWalker(node.Left)
		treeWalker(node.Right)
	}

	treeWalker(root)

	return result
}

func main() {
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{4, nil, nil}
	three := &TreeNode{2, nil, nil}
	four := &TreeNode{6, nil, nil}
	five := &TreeNode{7, nil, nil}
	six := &TreeNode{5, nil, nil}
	seven := &TreeNode{3, nil, nil}
	eight := &TreeNode{9, nil, nil}

	one.Left = two
	one.Right =  three

	two.Left = four
	two.Right = five

	three.Left = six
	three.Right = seven

	four.Left = eight

	preorderTraversal(one)
}