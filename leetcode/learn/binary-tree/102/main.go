package main

type TreeNode struct {
	Val         int
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

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	q := Queue{}
	q.Enqueue(root)

	for !q.Empty() {
		l := q.Length()

		var level []int
		for i := 0; i < l; i++ {
			node := q.Dequeue()

			if node.Left != nil {
				q.Enqueue(node.Left)
			}

			if node.Right != nil {
				q.Enqueue(node.Right)
			}
			level = append(level, node.Val)
		}
		result = append(result, level)
	}

	return result
}
