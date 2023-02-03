package main

func search(nums []int, target int) int {
	// Find rotation index
	rotationIndex := len(nums) - 1

	l := 0
	r := rotationIndex

	for l < r {
		pivot := l + (r-l)/2
		if nums[pivot] > nums[pivot+1] {
			rotationIndex = pivot
			break
		}

		if nums[pivot] > nums[l] {
			l = pivot + 1
		} else {
			r = pivot
		}
	}

	var start, end int
	if rotationIndex == len(nums)-1 {
		start = 0
		end = rotationIndex
	} else if target > nums[len(nums)-1] {
		start = 0
		end = rotationIndex
	} else {
		start = rotationIndex + 1
		end = len(nums) - 1
	}

	for start <= end {
		pivot := start + (end-start)/2

		if nums[pivot] == target {
			return pivot
		}

		if nums[pivot] < target {
			start = pivot + 1
		} else {
			end = pivot - 1
		}
	}

	return -1
}

func search1(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		pivot := l + (r-l)/2

		if nums[pivot] == target {
			return pivot
		}

		if nums[pivot] >= nums[l] {
			if target >= nums[l] && target <= nums[pivot] {
				r = pivot - 1
			} else {
				l = pivot + 1
			}
		} else {
			if target <= nums[r] && target >= nums[pivot] {
				l = pivot + 1
			} else {
				r = pivot - 1
			}
		}
	}

	return -1
}
