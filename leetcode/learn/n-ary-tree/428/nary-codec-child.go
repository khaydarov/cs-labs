package main

import (
	"fmt"
	"strconv"
	"strings"
)

func serialize1(root *Node) string {
	if root == nil {
		return ""
	}

	var result strings.Builder
	var helper func(node *Node)
	helper = func(node *Node) {
		if node == nil {
			return
		}

		result.WriteString(fmt.Sprintf("%d:%d;", node.Val, len(node.Children)))
		for _, child := range node.Children {
			helper(child)
		}
	}

	helper(root)
	return result.String()
}

func deserialize1(data string) *Node {
	if data == "" {
		return nil
	}

	nodes := strings.Split(data, ";")
	q := NewQueue(nodes)

	var helper func() *Node
	helper = func() *Node {
		if q.IsEmpty() {
			return nil
		}

		nodeStr, _ := q.Dequeue()
		nodeParts := strings.Split(nodeStr, ":")
		val, _ := strconv.Atoi(nodeParts[0])
		childrenCount, _ := strconv.Atoi(nodeParts[1])

		node := &Node{Val: val}
		for i := 0; i < childrenCount; i++ {
			node.Children = append(node.Children, helper())
		}

		return node
	}

	return helper()
}

func NewNaryCodec1() *Codec {
	return &Codec{
		serializer:   serialize1,
		deserializer: deserialize1,
	}
}
