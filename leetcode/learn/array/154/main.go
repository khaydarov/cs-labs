package main

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		if nums[r] > nums[l] {
			return nums[l]
		}

		m := (l + r) / 2

		if nums[l] < nums[m] {
			l = m + 1
		} else if nums[l] == nums[m] {
			l++
		} else {
			r = m
		}
	}

	return nums[l]
}
