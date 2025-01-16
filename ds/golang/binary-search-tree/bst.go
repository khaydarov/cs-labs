package main

import "fmt"

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

// Insert adds to tree a new node
// TC: O(h), where h is the height of the tree
// SC: O(h)
func (bst *BinarySearchTree) Insert(node *TreeNode) {
	var helper func (subtree, node *TreeNode) *TreeNode

	helper = func (subtree, node *TreeNode) *TreeNode {
		if subtree == nil {
			return node
		}

		if subtree.Val > node.Val {
			subtree.Left = helper(subtree.Left, node)
		} else {
			subtree.Right = helper(subtree.Right, node)
		}

		return subtree
	}

	helper(bst.root, node)
}

// Search lookups and returns a Node with passed value
// TC: O(h)
// SC: O(h)
func (bst *BinarySearchTree) Search(val int) *TreeNode {
	var helper func (subtree *TreeNode, val int) *TreeNode

	helper = func (subtree *TreeNode, target int) *TreeNode {
		if subtree == nil || subtree.Val == target {
			return subtree
		}

		if subtree.Val > target {
			return helper(subtree.Left, target)
		}

		return helper(subtree.Right, target)
	}

	return helper(bst.root, val)
}

// Delete removes from BST given value
func (bst *BinarySearchTree) Delete(val int) {
	var minValue func (node *TreeNode) *TreeNode
	minValue = func(node *TreeNode) *TreeNode {
		for node.Left != nil {
			node = node.Left
		}

		return node
	}

	target := bst.Search(val)
	if target == nil {
		return
	}

	var helper func (subtree *TreeNode, node *TreeNode) *TreeNode
	helper = func(subtree *TreeNode, node *TreeNode) *TreeNode {
		if subtree == nil {
			return nil
		}

		if subtree.Val > node.Val {
			subtree.Left = helper(subtree.Left, node)
		} else if subtree.Val < node.Val {
			subtree.Right = helper(subtree.Right, node)
		} else {
			if subtree.Left == nil && subtree.Right == nil {
				return nil
			}

			if subtree.Left != nil && subtree.Right == nil {
				return subtree.Left
			}

			if subtree.Right != nil && subtree.Left == nil {
				return subtree.Right
			}

			if subtree.Right != nil {
				inOrderSuccessor := minValue(subtree.Right)
				subtree.Val = inOrderSuccessor.Val

				subtree.Right = helper(subtree.Right, inOrderSuccessor)
			}
		}

		return subtree
	}

	helper(bst.root, target)
}

func main() {
	zero := &TreeNode{4, nil, nil}
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{6, nil, nil}
	three := &TreeNode{5, nil, nil}
	four := &TreeNode{8, nil, nil}

	zero.Left = one
	zero.Right = two

	two.Left = three
	two.Right = four

	bst := BinarySearchTree{zero}
	//bst.Insert(&TreeNode{Val: 2})
	bst.Delete(6)

	fmt.Println(bst.root.Right)
}