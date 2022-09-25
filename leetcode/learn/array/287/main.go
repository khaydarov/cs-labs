package main

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
