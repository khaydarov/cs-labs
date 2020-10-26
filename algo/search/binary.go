package search

// Template 1
func searchT1(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

// Template 2
func searchT2(nums []int, target int) int {
	l := 0
	r := len(nums)

	for l < r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}

	return -1
}

// Template 3
func searchT3(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l + 1 < r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			l = m
		} else {
			r = m
		}
	}

	if nums[l] == target {
		return l
	}

	if nums[r] == target {
		return r
	}

	return -1
}
