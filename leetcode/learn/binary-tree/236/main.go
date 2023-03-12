package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

		if l+r+m >= 2 {
			lca = root
		}

		return l+r+m > 0
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
	top := s.data[l-1]
	s.data = s.data[:l-1]

	return top
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

type Set struct {
	data    []*TreeNode
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

// TC: O(N)
// SC: O(N)
func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	var order []*TreeNode
	var first map[*TreeNode]int
	var height map[*TreeNode]int

	first = make(map[*TreeNode]int)
	height = make(map[*TreeNode]int)

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		order = append(order, node)
		height[node] = depth

		if node.Left != nil {
			dfs(node.Left, depth+1)
			order = append(order, node)
		}

		if node.Right != nil {
			dfs(node.Right, depth+1)
			order = append(order, node)
		}
	}

	dfs(root, 0)

	for i, v := range order {
		if _, ok := first[v]; !ok {
			first[v] = i
		}
	}

	l := first[p]
	r := first[q]

	if l > r {
		l, r = r, l
	}

	minNode := order[l]
	minHeight := height[minNode]
	for i := l + 1; i <= r; i++ {
		node := order[i]
		if minHeight > height[node] {
			minHeight = height[node]
			minNode = node
		}
	}

	return minNode
}
