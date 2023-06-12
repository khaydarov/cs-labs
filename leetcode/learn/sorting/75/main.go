package main

// TC: O(N)
// SC: O(1)
func sortColors(nums []int) {
	var zeros, ones, twos int
	for _, v := range nums {
		if v == 0 {
			zeros++
		} else if v == 1 {
			ones++
		} else {
			twos++
		}
	}

	pointer := 0
	for zeros+ones+twos > 0 {
		if zeros > 0 {
			nums[pointer] = 0
			zeros--
		} else if ones > 0 {
			nums[pointer] = 1
			ones--
		} else {
			nums[pointer] = 2
			twos--
		}

		pointer++
	}
}

// TC: O(N^2)
// SC: O(1)
func sortColors1(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		foundMinPosition := i
		for j := i + 1; j < n; j++ {
			if nums[foundMinPosition] > nums[j] {
				foundMinPosition = j
			}
		}

		nums[i], nums[foundMinPosition] = nums[foundMinPosition], nums[i]
	}
}

// TC: O(N)
// SC: O(1)
func sortColors2(nums []int) {
	var p0, p2, current int
	current = 0
	p0 = 0
	p2 = len(nums) - 1

	for current <= p2 {
		if nums[current] == 0 {
			nums[p0], nums[current] = nums[current], nums[p0]
			p0++
			current++
		} else if nums[current] == 2 {
			nums[p2], nums[current] = nums[current], nums[p2]
			p2--
		} else {
			current++
		}

	}
}
