package main

import (
	"sort"
)

func findDuplicate(nums []int) int {
	low := 1
	high := len(nums)

	answer := -1
	for low <= high {
		mid := low + (high-low)/2

		if countElementsSmallerOrEqualThan(nums, mid) > mid {
			answer = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return answer
}

func countElementsSmallerOrEqualThan(nums []int, target int) int {
	count := 0
	for _, v := range nums {
		if v <= target {
			count++
		}
	}

	return count
}

// Approach
//
// [1,2,3,4,4]
// [2,1,3,4,4]
// [3,1,2,4,4]
// [4,1,2,3,4]

// TC: O(N)
// SC: O(1)
func findDuplicate5(nums []int) int {
	for nums[0] != nums[nums[0]] {
		nums[0], nums[nums[0]] = nums[nums[0]], nums[0]
	}

	return nums[0]
}

// Approach
//
// [1,3,4,2,1]
// [3,1,4,2,1]
// [2,1,4,3,1]
// [4,1,2,3,1]
// [1,1,2,3,4]

// TC: O(N) - each number needs to be swapped at most once before it is places in its desired position
// SC: O(N)
func findDuplicate4(nums []int) int {
	duplicate := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		duplicate[i+1] = nums[i]
	}

	for i := 0; i < len(duplicate); i++ {
		for duplicate[i] != i {
			if duplicate[i] == duplicate[duplicate[i]] {
				return duplicate[i]
			}
			duplicate[i], duplicate[duplicate[i]] = duplicate[duplicate[i]], duplicate[i]
		}
	}

	return -1
}

// Approach: flip elements in position of nums[i]; if some element is already flipped then we found answer
//
// [1 2 3 4 4]
// [1 -2 3 4 4]
// [1 -2 -3 4 4]
// [1 -2 -3 -4 4]
// [1 -2 -3 -4 -4]
//
// TC: O(N)
// SC: O(1)
func findDuplicate3(nums []int) int {
	for i := 0; i < len(nums); i++ {
		current := abs(nums[i])
		if nums[current] < 0 {
			return current
		}

		nums[current] *= -1
	}

	return -1
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

// TC: O(N log N) - sorting; + O(N) - search;
// SC: O(log N) - sorting
func findDuplicate2(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}

	return -1
}

func findDuplicate1(nums []int) int {
	tortoise := nums[0]
	hare := nums[0]

	for true {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]

		if tortoise == hare {
			break
		}
	}

	tortoise = nums[0]
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[hare]
	}

	return hare
}

func main() {
	nums := []int{1, 3, 4, 2, 1}
	findDuplicate(nums)
}
