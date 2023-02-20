package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	start := find(nums, target, false)
	if start == -1 {
		return []int{-1, -1}
	}

	end := find(nums, target, true)

	return []int{start, end}
}

func find(nums []int, target int, right bool) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		pivot := l + (r-l)/2
		if nums[pivot] == target {
			if right && nums[pivot] < nums[pivot+1] {
				return pivot
			}

			if right {
				l = pivot + 1
			} else {
				r = pivot
			}
		} else if nums[pivot] < target {
			l = pivot + 1
		} else {
			r = pivot
		}
	}

	if l < len(nums) && nums[l] != target {
		return -1
	}

	return l
}
