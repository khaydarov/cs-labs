package main

import "fmt"

func main() {
	nums := []int{8, 1, 2, 2, 3}
	smallerNumbersThanCurrent(nums)
}

/**
 * TC = O (N * log N)
 * SC = O (N)
 */
func smallerNumbersThanCurrent(nums []int) []int {
	max := nums[0]
	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	max++

	all := make([]int, max)
	for _, v := range nums {
		all[v]++
	}

	sum := 0
	for i := 0; i < max; i++ {
		t := all[i]
		all[i] = sum
		sum += t
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = all[nums[i]]
	}

	fmt.Println(nums)
	return nums
}