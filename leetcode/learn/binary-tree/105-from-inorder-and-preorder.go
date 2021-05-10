package main

import "fmt"

type TreeNode struct {
	Val 		int
	Left, Right *TreeNode
}

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(x int) {
	q.data = append(q.data, x)
}

func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		return -1
	}

	front := q.data[0]
	q.data = q.data[1:]

	return front
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	queue := Queue{}

	for _, v := range preorder {
		queue.Enqueue(v)
	}

	var buildNode func(inorder []int) *TreeNode
	buildNode = func(inorder []int) *TreeNode {
		if len(inorder) == 0 || queue.Empty() {
			return nil
		}

		front := queue.Dequeue()
		divider := findInArray(inorder, front)

		root := &TreeNode{Val: front}
		root.Left = buildNode(inorder[:divider])
		root.Right = buildNode(inorder[divider + 1:])

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

func main() {
	r := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	//r := buildTree([]int{-1}, []int{-1})
	fmt.Println(r)
}