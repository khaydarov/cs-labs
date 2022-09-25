package main

import "fmt"

type TreeNode struct {
	Val 	int
	Cnt 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

type KthLargest struct {
	k 		int
	root 	*TreeNode
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val, Cnt: 1}
	}

	root.Cnt++
	if root.Val > val {
		root.Left = Insert(root.Left, val)
	} else {
		root.Right = Insert(root.Right, val)
	}

	return root
}

func Search(node *TreeNode, k int) *TreeNode {
	right := 0
	if node.Right != nil {
		right = node.Right.Cnt
	}

	if right + 1 == k {
		return node
	}

	if node.Right != nil && node.Right.Cnt < k {
		return Search(node.Left, k - node.Right.Cnt - 1)
	} else if node.Right != nil {
		return Search(node.Right, k)
	} else {
		return Search(node.Left, k - 1)
	}
}

func Constructor(k int, nums []int) KthLargest {
	var root *TreeNode
	for i := 0; i < len(nums); i++ {
		root = Insert(root, nums[i])
	}

	return KthLargest{k, root}
}

func (this *KthLargest) Add(val int) int {
	this.root = Insert(this.root, val)
	target := Search(this.root, this.k)

	return target.Val
}

func main() {
	//zero := &TreeNode{4, nil, nil}
	//one := &TreeNode{1, nil, nil}
	//two := &TreeNode{6, nil, nil}
	//three := &TreeNode{5, nil, nil}
	//four := &TreeNode{8, nil, nil}
	//
	//zero.Left = one
	//zero.Right = two
	//
	//two.Left = three
	//two.Right = four
	//
	//bst := BinarySearchTree{zero}
	////bst.Insert(&TreeNode{Val: 2})
	//bst.Delete(6)
	//
	//fmt.Println(bst.root.Right)
	r := Constructor(3, []int{0, 1})
	fmt.Println(r.Add(-1))
	fmt.Println(r.Add(1))
	fmt.Println(r.Add(-2))
	fmt.Println(r.Add(-4))
	fmt.Println(r.Add(3))
}