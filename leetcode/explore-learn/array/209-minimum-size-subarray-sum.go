package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	left := 0
	result := 100000
	for i, v := range nums {
		sum += v

		//[2, 3, 1, 2, 4, 3]
		for sum >= target {
			result = min(result, i - left + 1)
			sum -= nums[left]
			left++
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
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(minSubArrayLen(6, arr))
}
