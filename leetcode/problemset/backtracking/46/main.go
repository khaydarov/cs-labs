package main

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
