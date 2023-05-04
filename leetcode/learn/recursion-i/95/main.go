package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func serialize(node *TreeNode) string {
	var helper func(node *TreeNode) string
	helper = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}

		return strconv.Itoa(node.Val) + helper(node.Left) + helper(node.Right)
	}

	return helper(node)
}

func permute(nums []int) [][]int {
	var permutations [][]int
	var helper func(nums []int, start int)
	helper = func(nums []int, start int) {
		if start >= len(nums) {
			copyNums := make([]int, len(nums))
			copy(copyNums, nums)

			permutations = append(permutations, copyNums)
		}

		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			helper(nums, start+1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	helper(nums, 0)

	return permutations
}

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var helper func(node *TreeNode, val int) *TreeNode
	helper = func(node *TreeNode, val int) *TreeNode {
		if node == nil {
			return &TreeNode{Val: val}
		}

		if node.Val > val {
			node.Left = helper(node.Left, val)
		} else {
			node.Right = helper(node.Right, val)
		}

		return node
	}

	var node *TreeNode
	for _, v := range nums {
		node = helper(node, v)
	}

	return node
}

func generateTrees(n int) []*TreeNode {
	var result []*TreeNode

	var nums []int
	nums = make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	cache := make(map[string]bool)
	permutations := permute(nums)
	for _, nums := range permutations {
		node := buildTree(nums)
		serialized := serialize(node)
		if _, ok := cache[serialized]; ok {
			continue
		}

		cache[serialized] = true
		result = append(result, node)
	}

	return result
}
