package main

import "fmt"

func main() {
	cases := [][]int{
		{
			1, 3, 4, 2, 2,
		},
		{
			3, 1, 3, 4, 2,
		},
		{
			1, 1,
		},
		{
			1, 1, 2,
		},
		{
			1, 2, 2,
		},
		{
			2, 2, 2, 2, 2,
		},
	}

	for _, v := range cases {
		ans := findDuplicate(v)
		fmt.Println("answer is", ans)
	}
}

func findDuplicate(nums []int) int {
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