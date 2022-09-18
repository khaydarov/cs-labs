package main

import "fmt"

func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		m := (l + r) / 2
		if nums[m] >= nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	fmt.Println(l)

	return l
}
