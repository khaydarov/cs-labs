package main

import (
	"fmt"
	"math"
)

func splitArray(nums []int, k int) int {
	minSum := math.MaxUint32
	var helper func(arr []int, c, k, maxSum int)

	helper = func(arr []int, c, k, maxSum int) {
		//if k == 0 {
		//	currentSum := 0
		//	for i := c; i < len(arr); i++ {
		//		currentSum += arr[i]
		//	}
		//
		//	if currentSum > maxSum {
		//		maxSum = currentSum
		//	}
		//
		//	if minSum > maxSum {
		//		minSum = maxSum
		//	}
		//}

		currentSum := 0
		for i := c; i < len(arr); i++ {
			currentSum += arr[i]
			if currentSum > maxSum {
				maxSum = currentSum
			}
			//helper(arr, i+1, k-1, maxSum)
		}
	}

	helper(nums, 0, k-1, 0)
	return minSum
}

func main() {
	nums := []int{7, 2, 5, 10, 8}
	//nums := []int{1, 2, 3, 4, 5}
	k := 2
	r := splitArray(nums, k)
	fmt.Println(r)
}
