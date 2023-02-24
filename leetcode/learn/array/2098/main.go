package main

import "sort"

// Approach 1
// [6,1,5,7,3,1]
// [1,1,3,5,6,7]
// 1) Sort array
// 2) Sum last k - 1 elements and find third that gives even sum
// TC: O(N log N + N)
// SC: O(1)

// Approach 2
// 2,0,5,3,4
// 0,2,3,4,5
// minOdd = 5
// minEven = -1
// maxOdd = 3
// maxEven = 4
func largestEvenSum(nums []int, k int) int64 {
	sort.Ints(nums)

	var total int64
	minEven := -1
	minOdd := -1
	for i := len(nums) - 1; i >= len(nums)-k; i-- {
		total += int64(nums[i])
		if nums[i]%2 == 0 {
			minEven = nums[i]
		} else {
			minOdd = nums[i]
		}
	}

	if total%2 == 0 {
		return total
	}

	maxEven := -1
	maxOdd := -1
	for i := 0; i < len(nums)-k; i++ {
		if nums[i]%2 == 0 {
			maxEven = max(maxEven, nums[i])
		} else {
			maxOdd = max(maxOdd, nums[i])
		}
	}

	var result int64
	result = -1

	if minEven != -1 && maxOdd != -1 {
		if result < total-int64(minEven)+int64(maxOdd) {
			result = total - int64(minEven) + int64(maxOdd)
		}
	}

	if minOdd != -1 && maxEven != -1 {
		if result < total-int64(minOdd)+int64(maxEven) {
			result = total - int64(minOdd) + int64(maxEven)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}