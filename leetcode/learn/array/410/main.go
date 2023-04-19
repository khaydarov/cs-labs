package main

import "fmt"

// Top-Down dynamic programming approach
//
// 7,2,5,10,8
// [7 - 2 23] - 23
// [7 - 7 18] - 18
// [7 - 17 8] - 17
// cache[1:2] = 17
// [9 - 5 18] - 18
// [9 - 15 8] - 15
// cache[2:2] = 15
// [14 - 10 8] - 10
// cache[3:2] = 10
func splitArray(nums []int, k int) int {
	prefixSum := retrievePrefixSum(nums)
	cache := make(map[string]int)

	var getMinimumLargestSplitSum func(currentIndex int, subArrayCount int) int
	getMinimumLargestSplitSum = func(currentIndex int, subArrayCount int) int {
		key := fmt.Sprintf("%d:%d", currentIndex, subArrayCount)

		if v, ok := cache[key]; ok {
			return v
		}

		if subArrayCount == 1 {
			return prefixSum[len(nums)-1] - prefixSum[currentIndex] + nums[currentIndex]
		}

		minimumLargestSplitSum := prefixSum[len(nums)-1]
		currentSum := 0
		for i := currentIndex; i <= len(nums)-subArrayCount; i++ {
			currentSum += nums[i]
			largestSplitSum := max(
				currentSum,
				getMinimumLargestSplitSum(i+1, subArrayCount-1),
			)

			minimumLargestSplitSum = min(minimumLargestSplitSum, largestSplitSum)
		}

		cache[key] = minimumLargestSplitSum

		return minimumLargestSplitSum
	}

	return getMinimumLargestSplitSum(0, k)
}

// Bottom-Up dynamic programming approach
// 7,2,5,10,8
// memo[0][1] = 32
// memo[1][1] = 25
// memo[2][1] = 23
// memo[3][1] = 18
// memo[4][1] = 8
//
// memo[currIndex][2] = max(sum[0..currIndex], memo[currIndex + 1][1])
func splitArray1(nums []int, k int) int {
	memo := make([][]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		memo[i] = make([]int, len(nums)+1)
	}

	prefixSum := retrievePrefixSum(nums)

	for subArrayCnt := 1; subArrayCnt <= k; subArrayCnt++ {
		for currIdx := 0; currIdx < len(nums); currIdx++ {
			if subArrayCnt == 1 {
				memo[currIdx][subArrayCnt] = prefixSum[len(nums)] - prefixSum[currIdx]
				continue
			}

			minLargestSum := prefixSum[len(nums)]
			for i := currIdx; i <= len(nums)-subArrayCnt; i++ {
				firstLargestSum := prefixSum[i+1] - prefixSum[currIdx]

				largestSum := max(
					firstLargestSum,
					memo[i+1][subArrayCnt-1],
				)

				minLargestSum = min(largestSum, minLargestSum)

				if firstLargestSum >= minLargestSum {
					break
				}
			}

			memo[currIdx][subArrayCnt] = minLargestSum
		}
	}

	return memo[0][k]
}

func retrievePrefixSum(nums []int) []int {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = nums[i] + prefixSum[i-1]
	}

	return prefixSum
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
