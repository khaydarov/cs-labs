package main

import (
	"strconv"
	"strings"
)

type BinaryNode struct {
	Val   int
	Left  *BinaryNode
	Right *BinaryNode
}

func serialize(root *BinaryNode) string {
	if root == nil {
		return "#"
	}

	return strconv.Itoa(root.Val) + "." + serialize(root.Left) + "." + serialize(root.Right)
}

func deserialize(data string) *BinaryNode {
	vals := strings.Split(data, ".")
	q := NewQueue(vals)

	var buildTree func() *BinaryNode
	buildTree = func() *BinaryNode {
		if q.IsEmpty() {
			return nil
		}

		front, err := q.Dequeue()
		if err != nil {
			return nil
		}

		if front == "#" {
			return nil
		}

		val, _ := strconv.Atoi(front)

		node := &BinaryNode{Val: val}
		node.Left = buildTree()
		node.Right = buildTree()

		return node
	}

	return buildTree()
}

func convertToTreeNode(root *Node) *BinaryNode {
	if root == nil {
		return nil
	}

	treeNode := &BinaryNode{Val: root.Val}
	if len(root.Children) > 0 {
		firstNode := root.Children[0]
		treeNode.Left = convertToTreeNode(firstNode)
	}

	sibling := treeNode.Left
	for i := 1; i < len(root.Children); i++ {
		sibling.Right = convertToTreeNode(root.Children[i])
		sibling = sibling.Right
	}

	return treeNode
}

func convertToNode(root *BinaryNode) *Node {
	if root == nil {
		return nil
	}

	node := &Node{Val: root.Val}
	sibling := root.Left
	for sibling != nil {
		node.Children = append(node.Children, convertToNode(sibling))
		sibling = sibling.Right
	}

	return node
}

func NewBinaryCodec() *Codec {
	return &Codec{
		serializer: func(root *Node) string {
			return serialize(convertToTreeNode(root))
		},
		deserializer: func(data string) *Node {
			return convertToNode(deserialize(data))
		},
	}
}
