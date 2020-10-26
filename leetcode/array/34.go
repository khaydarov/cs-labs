package main

import "fmt"

func main() {
	nums := []int{1, 2, 2}
	r := searchRange(nums, 2)
	fmt.Println(r)
}

func searchRange(nums []int, target int) []int {
	l := binSearch(nums, target, false)

	if len(nums) == 0 || nums[l] != target {
		return []int{-1, -1}
	}

	r := binSearch(nums, target, true)

	if nums[r] == target {
		return []int{l, r}
	} else {
		return []int{l, r - 1}
	}
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

	return l
}