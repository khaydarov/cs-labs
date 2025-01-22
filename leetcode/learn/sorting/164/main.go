package main

const NUM_DIGITS = 10

func countingSort(nums []int, placeVal int) {
	n := len(nums)
	counts := make([]int, NUM_DIGITS)
	for i := 0; i < n; i++ {
		current := nums[i] / placeVal
		counts[current%NUM_DIGITS]++
	}

	var count, startingIndex int
	for i := 0; i < NUM_DIGITS; i++ {
		count = counts[i]
		counts[i] = startingIndex
		startingIndex += count
	}

	sorted := make([]int, n)
	for i := 0; i < n; i++ {
		current := nums[i] / placeVal
		sorted[counts[current%NUM_DIGITS]] = nums[i]
		counts[current%NUM_DIGITS]++
	}

	for i := 0; i < n; i++ {
		nums[i] = sorted[i]
	}
}

func RadixSort(nums []int) {
	maxValue := nums[0]
	for _, v := range nums {
		if maxValue < v {
			maxValue = v
		}
	}

	placeVal := 1
	for placeVal <= maxValue {
		countingSort(nums, placeVal)
		placeVal *= 10
	}
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	RadixSort(nums)
	maxDiff := nums[1] - nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > maxDiff {
			maxDiff = nums[i] - nums[i-1]
		}
	}

	return maxDiff
}
