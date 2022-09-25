package main

import "fmt"

/**
 * O(N)
 */
func minSubArrayLen(target int, nums []int) int {
	sum := 0
	left := 0
	result := 100000
	for i, v := range nums {
		sum += v

		//[2, 3, 1, 2, 4, 3]
		for sum >= target {
			result = min(result, i-left+1)
			sum -= nums[left]
			left++
		}
	}

	if result == 100000 {
		return 0
	}

	return result
}

/**
 * TC: O(NlogN)
 */
func minSubArrayLen1(target int, nums []int) int {
	result := 100000
	sums := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]
	}

	for i := 0; i < len(sums)-1; i++ {
		l := i + 1
		r := len(sums) - 1

		for l <= r {
			m := (l + r) / 2
			if sums[m]-sums[i] >= target {
				result = min(result, m-i)
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	if result == 100000 {
		return 0
	}

	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func main() {
	arr := []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}
	fmt.Println(minSubArrayLen(213, arr))
}
