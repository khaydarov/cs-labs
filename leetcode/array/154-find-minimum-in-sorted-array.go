package main

import "fmt"

func main() {
	cases := [][]int{
		{
			1, 2, 3, 4, 7, 7, 7, 7, 7, 7,
		},
		{
			1, 2, 3, 4, 7, 7, 7, 7, 7, 7, 0,
		},
		{
			1, 3, 5,
		},
		{
			2, 2, 2, 2, 0, 1,
		},
		{
			4, 5, 5, 5, 5, 0, 1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4,
		},
		{
			3, 3, 1, 3,
		},
		{
			1, 1, 1,
		},
		{
			10, 1, 10, 10, 10,
		},
	}

	for _, v := range cases {
		ans := findMin(v)
		fmt.Println("answer is", ans)
	}
}

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		m := (l + r) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else if nums[m] < nums[r] {
			r = m
		} else {
			r--
		}
	}

	return nums[l]
}