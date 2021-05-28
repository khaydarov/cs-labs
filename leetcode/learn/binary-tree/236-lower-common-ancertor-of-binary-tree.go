package main

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

// TC: O(N)
// SC: O(N)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var left, right *TreeNode
	var dfs func(current *TreeNode)

	dfs = func(current *TreeNode) {
		if current == nil {
			return
		}

		if left != nil && right == nil && (current == p || current == q) {
			right = current
		}

		if left == nil && (current == p || current == q) {
			left = current
		}

		dfs(current.Left)

		if left != nil && right == nil && left == current.Left {
			left = current
		}

		dfs(current.Right)

		if left != nil && right == nil && left == current.Right {
			left = current
		}
	}

	dfs(root)

	return left
}

// TC: O(N)
// SC: O(N)
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode
	var traverser func(root *TreeNode) bool
	
	traverser = func(root *TreeNode) bool {
		if root == nil {
			return false
		}

		var l, r, m int

		left := traverser(root.Left)
		right := traverser(root.Right)

		if left == true {
			l = 1
		}

		if right == true {
			r = 1
		}

		if root == p || root == q {
			m = 1
		}

		if l + r + m >= 2 {
			lca = root
		}

		return l + r + m > 0
	}

	traverser(root)

	return lca
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

	l := len(s.data)
	top := s.data[l - 1]
	s.data = s.data[:l - 1]

	return top
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

type Set struct {
	data []*TreeNode
	hashMap map[*TreeNode]bool
}

func (s *Set) Add(node *TreeNode) {
	if s.hashMap[node] {
		return
	}

	s.data = append(s.data, node)
	s.hashMap[node] = true
}

func (s *Set) Contains(node *TreeNode) bool {
	return s.hashMap[node]
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	parent := make(map[*TreeNode]*TreeNode)

	stack := Stack{}
	stack.Push(root)
	parent[root] = nil

	for !stack.Empty() {
		current := stack.Pop()

		if current.Left != nil {
			stack.Push(current.Left)
			parent[current.Left] = current
		}

		if current.Right != nil {
			stack.Push(current.Right)
			parent[current.Right] = current
		}
	}

	ancestors := Set{
		hashMap: make(map[*TreeNode]bool),
	}

	for p != nil {
		ancestors.Add(p)
		p = parent[p]
	}

	for !ancestors.Contains(q) {
		q = parent[q]
	}

	return q
}

func main() {
	zero := &TreeNode{0, nil, nil}
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{2, nil, nil}
	three := &TreeNode{3, nil, nil}
	four := &TreeNode{4, nil, nil}
	five := &TreeNode{5, nil, nil}
	six := &TreeNode{6, nil, nil}
	seven := &TreeNode{7, nil, nil}

	zero.Left = one

	one.Left = two
	one.Right = three

	two.Left = four
	two.Right = five

	three.Left = six
	three.Right = seven
	//fmt.Println(lowestCommonAncestor(zero, two, three))
	lowestCommonAncestor2(nil, two, three)
}

