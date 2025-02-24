package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Queue []string

func NewQueue(vals []string) *Queue {
	q := Queue{}
	for _, val := range vals {
		q.Enqueue(val)
	}

	return &q
}

func (q *Queue) Enqueue(val string) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue() (string, error) {
	if len(*q) == 0 {
		return "", errors.New("queue is empty")
	}

	val := (*q)[0]
	*q = (*q)[1:]
	return val, nil
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

type Node struct {
	Val      int
	Children []*Node
}

func serialize(root *Node) string {
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

func deserialize(data string) *Node {
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

func main() {
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}
	four := &Node{Val: 4}
	five := &Node{Val: 5}

	one.Children = []*Node{two, three}
	two.Children = []*Node{four}
	four.Children = []*Node{five}

	r := serialize(one)
	fmt.Println("serialized", r)
	node := deserialize(r)

	if serialize(node) == r {
		fmt.Println("check OK")
	} else {
		fmt.Println("check FAIL")
	}
}
