package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// TC: O(N)
// SC: O(N)
func sortedArrayToBST(nums []int) *TreeNode {
	var helper func (root *TreeNode, nums []int, low, high int) *TreeNode
	helper = func (root *TreeNode, nums []int, low, high int) *TreeNode {
		mid := (low + high) / 2

		if root == nil {
			root = &TreeNode{Val: nums[mid]}
		}

		if mid - 1 >= low {
			root.Left = helper(nil, nums, low, mid)
		}

		if mid + 1 <= high - 1 {
			root.Right = helper(nil, nums, mid + 1, high)
		}

		return root
	}

	return helper(nil, nums, 0, len(nums))
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5}
	sortedArrayToBST(nums)
}