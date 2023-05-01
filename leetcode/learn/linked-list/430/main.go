package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

// TC: O(N^2)
// SC: O(N)

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	current := &Node{}
	head := current
	for root != nil {
		current.Next = root
		root.Prev = current
		current = current.Next

		if root.Child != nil {
			next := root.Next

			rest := flatten(root.Child)
			current.Next = rest
			rest.Prev = current

			for current.Next != nil {
				current = current.Next
			}

			root.Child = nil
			root = next
		} else {
			root = root.Next
		}
	}

	head.Next.Prev = nil
	return head.Next
}

// TC: O(N)
// SC: O(N)
func flatten1(root *Node) *Node {
	var _flatten func(node *Node) *Node

	// returns last node
	_flatten = func(node *Node) *Node {
		if node == nil {
			return node
		}

		current := node
		previous := current
		for current != nil {
			if current.Child == nil {
				previous = current
				current = current.Next

				continue
			}

			next := current.Next
			child := current.Child

			lastOfRest := _flatten(child)
			current.Child = nil

			current.Next = child
			child.Prev = current

			lastOfRest.Next = next

			if next != nil {
				next.Prev = lastOfRest
			}

			previous = lastOfRest
			current = next
		}

		return previous
	}

	_flatten(root)

	return root
}
