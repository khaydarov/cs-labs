package main

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		pivot := l + (r-l)/2

		if nums[pivot] < nums[r] {
			r = pivot
		} else if nums[pivot] > nums[r] {
			l = pivot + 1
		} else {
			r--
		}
	}

	return nums[l]
}
