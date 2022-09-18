package main

func searchRange(nums []int, target int) []int {
	l := binSearch(nums, target, false)

	if len(nums) == 0 || nums[l] != target {
		return []int{-1, -1}
	}
	r := binSearch(nums, target, true)
	return []int{l, r}
}

func binSearch(arr []int, target int, right bool) int {
	l := 0
	r := len(arr) - 1

	for l < r {
		m := (l + r) / 2

		if arr[m] < target || (right && arr[m] == target) {
			l = m + 1
		} else {
			r = m
		}
	}

	if arr[r] == target {
		return r
	} else {
		return r - 1
	}
}
