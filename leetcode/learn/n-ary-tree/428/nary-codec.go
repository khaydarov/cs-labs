package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Identifier int
type Counter Identifier

func (c *Counter) Inc() Identifier {
	*c++
	return Identifier(*c)
}

func (c *Counter) Val() Identifier {
	return Identifier(*c)
}

func serializer(root *Node) string {
	if root == nil {
		return ""
	}

	idGenerator := Counter(1)
	nodesStr := make([]string, 0)

	var helper func(node *Node, parentNode *Node, parentId Identifier)
	helper = func(node *Node, parentNode *Node, parentId Identifier) {
		if node == nil {
			return
		}

		nodeStr := fmt.Sprintf("%d:%d:%d", idGenerator.Val(), node.Val, parentId)
		nodesStr = append(nodesStr, nodeStr)

		parentId = idGenerator.Val()
		for _, child := range node.Children {
			idGenerator.Inc()
			helper(child, node, parentId)
		}
	}

	helper(root, nil, Identifier(1))
	return strings.Join(nodesStr, ";")
}

func deserializer(data string) *Node {
	if data == "" {
		return nil
	}

	hm := make(map[string]*Node)
	nodes := strings.Split(data, ";")

	for _, nodeStr := range nodes {
		parts := strings.Split(nodeStr, ":")
		val, _ := strconv.Atoi(parts[1])
		hm[parts[0]] = &Node{Val: val}
	}

	for _, nodeStr := range nodes {
		parts := strings.Split(nodeStr, ":")
		parentId := parts[2]

		if parts[0] == parentId {
			continue
		}

		parent := hm[parentId]
		parent.Children = append(parent.Children, hm[parts[0]])
	}

	return hm["1"]
}

func NewNaryCodec() *Codec {
	return &Codec{
		serializer:   serializer,
		deserializer: deserializer,
	}
}
