package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	data []*TreeNode
}

func (s *Stack) Push(node *TreeNode) {
	s.data = append(s.data, node)
}

func (s *Stack) Pop() *TreeNode {
	if len(s.data) == 0 {
		return nil
	}

	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return top
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

type BSTIterator struct {
	stack Stack
}

func Constructor(root *TreeNode) BSTIterator {
	stack := Stack{}

	head := root
	for head != nil {
		stack.Push(head)
		head = head.Left
	}

	return BSTIterator{stack}
}

func (this *BSTIterator) Next() int {
	top := this.stack.Pop()

	if top.Right != nil {
		head := top.Right
		for head != nil {
			this.stack.Push(head)
			head = head.Left
		}
	}

	return top.Val
}

func (this *BSTIterator) HasNext() bool {
	return !this.stack.Empty()
}

func main() {
	zero := &TreeNode{5, nil, nil}
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{4, nil, nil}
	three := &TreeNode{3, nil, nil}
	four := &TreeNode{6, nil, nil}

	zero.Left = one
	zero.Right = two

	two.Left = three
	two.Right = four

	iterator := Constructor(zero)
	iterator.HasNext()
	iterator.Next()
}
