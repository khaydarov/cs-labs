package main

func QuickSelect(nums []int, left, right, k int) int {
	pivot := partition(nums, left, right)

	if pivot == k {
		return nums[pivot]
	} else if pivot < k {
		return QuickSelect(nums, pivot+1, right, k)
	} else {
		return QuickSelect(nums, left, pivot-1, k)
	}
}

func partition(nums []int, left, right int) int {
	index := left + (right-left)/2
	nums[index], nums[right] = nums[right], nums[index]

	for i := left; i < right; i++ {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	return left
}
