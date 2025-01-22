package radix

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

// LSD (Least Significant Digit) Radix Sort
// TC: O(W*(N+K)) where W is the number of digits in the largest number and N is the number of elements
// SC: O(N+K)
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
