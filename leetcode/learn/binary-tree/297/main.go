package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LinkedList struct {
	Val  string
	Next *LinkedList
	Prev *LinkedList
}

type Queue struct {
	head *LinkedList
	tail *LinkedList
	size int
}

func (q *Queue) Enqueue(b string) {
	item := &LinkedList{Val: b}
	item.Next = q.tail
	item.Prev = q.tail.Prev

	q.tail.Prev.Next = item
	q.tail.Prev = item
	q.size++
}

func (q *Queue) Dequeue() string {
	if q.Size() == 0 {
		return ""
	}

	front := q.head.Next
	front.Prev.Next = front.Next
	front.Next.Prev = front.Prev

	q.size--

	return front.Val
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Println() {
	current := q.head.Next
	for current != q.tail {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func CreateQueue() Queue {
	head := &LinkedList{}
	tail := &LinkedList{}

	head.Next = tail
	tail.Prev = head

	return Queue{
		head,
		tail,
		0,
	}
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	q := CreateQueue()

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			q.Enqueue("#")
			return
		}

		q.Enqueue(strconv.Itoa(node.Val))

		helper(node.Left)
		helper(node.Right)
	}

	helper(root)

	var result string
	for q.Size() > 0 {
		result += q.Dequeue()
		result += "."
	}

	return result
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	q := CreateQueue()

	i := 0
	for i < len(data) {
		var val string
		for i < len(data) && data[i] != '.' {
			val += string(data[i])
			i++
		}

		q.Enqueue(val)
		i++
	}

	var treeBuilder func() *TreeNode
	treeBuilder = func() *TreeNode {
		if q.Size() == 0 {
			return nil
		}

		front := q.Dequeue()
		if front == "#" {
			return nil
		}

		val, _ := strconv.Atoi(front)

		node := &TreeNode{Val: val}
		node.Left = treeBuilder()
		node.Right = treeBuilder()

		return node
	}

	return treeBuilder()
}
