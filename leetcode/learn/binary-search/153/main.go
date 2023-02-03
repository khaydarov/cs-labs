package main

func findMin(nums []int) int {
	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}

	l := 0
	r := len(nums) - 1
	for l < r {
		pivot := l + (r-l)/2

		if nums[pivot] > nums[pivot+1] {
			return nums[pivot+1]
		}

		if nums[pivot] > nums[r] {
			l = pivot + 1
		} else {
			r = pivot
		}
	}

	return nums[l]
}
