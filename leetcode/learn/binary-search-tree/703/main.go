package main

type TreeNode struct {
	Val         int
	RightHeight int
	Left        *TreeNode
	Right       *TreeNode
}

type KthLargest struct {
	k    int
	root *TreeNode
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{
		k: k,
	}
	kthLargest.Init(nums)

	return kthLargest
}

func (this *KthLargest) insert(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{
			Val:         val,
			RightHeight: 1,
		}
	}

	if node.Val > val {
		node.Left = this.insert(node.Left, val)
	} else if node.Val <= val {
		node.RightHeight++
		node.Right = this.insert(node.Right, val)
	}

	return node
}

func (this *KthLargest) findKthLargest(node *TreeNode, k int) *TreeNode {
	if node.RightHeight == k {
		return node
	}

	if node.RightHeight > k {
		return this.findKthLargest(node.Right, k)
	}

	return this.findKthLargest(node.Left, k-node.RightHeight)
}

func (this *KthLargest) Add(val int) int {
	this.root = this.insert(this.root, val)

	target := this.findKthLargest(this.root, this.k)

	return target.Val
}

func (this *KthLargest) Init(nums []int) {
	for _, v := range nums {
		this.root = this.insert(this.root, v)
	}
}
