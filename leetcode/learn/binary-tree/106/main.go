package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	data []int
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}

	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return top
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

// TC: O(N^2)
// SC: O(N)
func buildTree(inorder []int, postorder []int) *TreeNode {
	stack := Stack{}

	for _, v := range postorder {
		stack.Push(v)
	}

	var buildNode func(inorder []int) *TreeNode
	buildNode = func(inorder []int) *TreeNode {
		if len(inorder) == 0 || stack.Empty() {
			return nil
		}

		top := stack.Pop()
		divider := findInArray(inorder, top)

		root := &TreeNode{Val: top}
		root.Right = buildNode(inorder[divider+1:])
		root.Left = buildNode(inorder[:divider])

		return root
	}

	return buildNode(inorder)
}

// TC: O(N)
// SC: O(1)
func findInArray(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}

	return -1
}
