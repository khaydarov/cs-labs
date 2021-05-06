package main

import "fmt"

type TreeNode struct {
	Val 		int
	Left, Right *TreeNode
}

type Queue struct {
	data []*TreeNode
}

func (q *Queue) Enqueue(node *TreeNode) {
	q.data = append(q.data, node)
}

func (q *Queue) Dequeue() *TreeNode {
	if len(q.data) == 0 {
		return nil
	}

	node := q.data[0]
	q.data = q.data[1:]

	return node
}

func (q *Queue) Length() int {
	return len(q.data)
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

// Using top-down approach
//
// TC: O(V) where V is the number of nodes
// SC: O(V)
func maxDepth1(root *TreeNode) int {
	var result int
	var traversal func (node *TreeNode, depth int)

	traversal = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			result = max(result, depth)
		}

		traversal(node.Left, depth + 1)
		traversal(node.Right, depth + 1)
	}

	traversal(root, 1)

	return result
}

// Using BFS
//
// TC: O(V)
// SC: O(V)
func maxDepth2(root *TreeNode) int {
	q := Queue{}
	q.Enqueue(root)

	steps := 0
	for !q.Empty() {
		l := q.Length()

		for i := 0; i < l; i++ {
			node := q.Dequeue()

			if node.Left != nil {
				q.Enqueue(node.Left)
			}

			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}

		steps++
	}

	fmt.Println(steps)
	return steps
}

// Using bottom-up approach
//
// TC: O(V)
// SC: O(V)
func maxDepth3(root *TreeNode) int {
	var traversal func (node *TreeNode, depth int) int
	traversal = func(node *TreeNode, depth int) int {
		if node == nil {
			return depth - 1
		}

		left := traversal(node.Left, depth + 1)
		right := traversal(node.Right, depth + 1)

		return max(left, right)
	}

	result := traversal(root, 1)

	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func main() {
	//one := &TreeNode{1, nil, nil}
	//two := &TreeNode{2, nil, nil}
	//three := &TreeNode{3, nil, nil}
	//four := &TreeNode{4, nil, nil}
	//five := &TreeNode{5, nil, nil}
	//six := &TreeNode{6, nil, nil}
	//seven := &TreeNode{7, nil, nil}
	//eight := &TreeNode{8, nil, nil}
	//
	//one.Left = two
	//one.Right = three
	//
	//two.Left = four
	//two.Right = five
	//
	//five.Left = six
	//five.Right = seven
	//
	//three.Right = eight

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

	maxDepth1(nil)
}