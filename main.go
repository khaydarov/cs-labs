package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	copies := make(map[int]*Node)

	clone := &Node{}
	_helper(node, clone, copies)

	return clone
}

func _helper(node *Node, clone *Node, copies map[int]*Node) {
	if node == nil {
		return
	}

	clone.Val = node.Val
	copies[node.Val] = clone

	for _, neighbor := range node.Neighbors {
		if v, ok := copies[neighbor.Val]; ok {
			clone.Neighbors = append(clone.Neighbors, v)
		} else {
			neighborCopy := &Node{}
			_helper(neighbor, neighborCopy, copies)

			clone.Neighbors = append(clone.Neighbors, neighborCopy)
		}
	}
}

func main() {
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}
	four := &Node{Val: 4}

	one.Neighbors = append(one.Neighbors, two)
	one.Neighbors = append(one.Neighbors, four)

	two.Neighbors = append(two.Neighbors, one)
	two.Neighbors = append(two.Neighbors, three)

	three.Neighbors = append(three.Neighbors, two)
	three.Neighbors = append(three.Neighbors, four)

	four.Neighbors = append(four.Neighbors, one)
	four.Neighbors = append(four.Neighbors, three)

	//r := cloneGraph(one)

	//return r
}
