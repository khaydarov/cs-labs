package main

type Node struct {
	Val    int
	Prev   *Node
	Next   *Node
	Random *Node
}

// TC: O(N)
// SC: O(1)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	current := head
	for current != nil {
		next := current.Next

		clone := &Node{Val: current.Val}
		current.Next = clone
		clone.Next = next

		current = next
	}

	current = head
	for current != nil {
		clone := current.Next

		if current.Random != nil {
			clone.Random = current.Random.Next
		}

		current = clone.Next
	}

	current = head
	clone := current.Next

	dummy := &Node{}
	dummy.Next = clone
	for current != nil {
		current.Next = clone.Next

		if current.Next != nil {
			clone.Next = current.Next.Next
		} else {
			clone.Next = nil
		}

		current = current.Next
		clone = clone.Next
	}

	return dummy.Next
}

// TC: O(N)
// SC: O(N)
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}

	m := make(map[*Node]*Node, 1000)

	current := head
	for current != nil {
		m[current] = &Node{Val: current.Val}
		current = current.Next
	}

	current = head
	copyHead := m[current]
	for current != nil {
		if _, ok := m[current]; !ok {
			continue
		}

		copyNode := m[current]
		copyNode.Next = m[current.Next]
		copyNode.Random = m[current.Random]

		current = current.Next
	}

	return copyHead
}

// TC: O(N)
// SC: O(N)
func copyRandomList2(head *Node) *Node {
	cache := make(map[*Node]*Node)
	var helper func(node *Node) *Node
	helper = func(node *Node) *Node {
		if node == nil {
			return node
		}

		if copyNode, ok := cache[node]; ok {
			return copyNode
		}

		copyNode := &Node{Val: node.Val}
		cache[node] = copyNode

		copyNode.Next = helper(node.Next)
		copyNode.Random = helper(node.Random)

		return copyNode
	}

	return helper(head)
}
