package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

type LinkedList struct {
	Val 	string
	Next 	*LinkedList
}

type Queue struct {
	data *LinkedList
}

func (q *Queue) Enqueue(x string) {
	if q.data == nil {
		q.data = &LinkedList{Val: x}
	} else {
		current := q.data
		for current.Next != nil {
			current = current.Next
		}

		current.Next = &LinkedList{Val: x}
	}
}

func (q *Queue) Dequeue() string {
	if q.data == nil {
		return ""
	}

	front := q.data.Val
	q.data = q.data.Next

	return front
}

func (q *Queue) Empty() bool {
	return q.data == nil
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var result string

	var traverser func (node *TreeNode)
	traverser = func (node *TreeNode) {
		if node == nil {
			result += "#"
			result += "."
			return
		}

		result += strconv.Itoa(node.Val)
		result += "."

		traverser(node.Left)
		traverser(node.Right)
	}

	traverser(root)

	return result
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}


	q := Queue{}
	i := 0
	for i < len(data) {
		k := ""
		for i < len(data) && data[i] != '.' {
			k += string(data[i])
			i++
		}

		q.Enqueue(k)
		i++
	}

	var traverser func (node *TreeNode) *TreeNode
	traverser = func (node *TreeNode) *TreeNode {
		if q.Empty() {
			return nil
		}

		front := q.Dequeue()
		if front == "#" {
			return nil
		}

		value, _ := strconv.Atoi(front)

		node.Val = value
		node.Left = traverser(&TreeNode{})
		node.Right = traverser(&TreeNode{})

		return node
	}


	root := traverser(&TreeNode{})

	return root
}

func main() {
	zero := &TreeNode{0, nil, nil}
	one := &TreeNode{-123, nil, nil}
	two := &TreeNode{-32, nil, nil}
	three := &TreeNode{31, nil, nil}
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

	codec := Constructor()
	r := codec.serialize(one)
	fmt.Println(r)

	d := codec.deserialize(r)
	fmt.Println(d, d.Left, d.Right)
}

