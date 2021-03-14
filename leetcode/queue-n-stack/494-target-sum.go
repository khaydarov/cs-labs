package main

import "fmt"

// Approach 1
// Brute force all solutions and increment counter
// TC: O(2^N)
// SC: O(N)
//
// Approach 2
// Optimize Approach 1: add memoization
// TC: O(l*N)
// SC: O(l*N), l - the range of sum, n - size of array
//
// Approach 3
//
func findTargetSumWays(nums []int, S int) int {
	var calculate func(nums []int, i, sum, S int) int

	cache := make(map[string]int)
	calculate = func(nums []int, i, sum, S int) int {
		if i == len(nums) {
			if sum == S {
				return 1
			} else {
				return 0
			}
		}

		key := fmt.Sprintf("%d:%d", i, sum)
		if v, ok := cache[key]; ok {
			return v
		}

		add := calculate(nums, i + 1, sum + nums[i], S)
		substract := calculate(nums, i + 1, sum - nums[i], S)

		cache[key] = add + substract

		return cache[key]
	}

	return calculate(nums, 0, 0, S)
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	r := findTargetSumWays(nums, 3)
	fmt.Println(r)
}