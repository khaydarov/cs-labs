package main

// Approach 1
// [1 2 2 2 5 0]; n = 5
// [1 3 5 7 12 12]
//

// TC: O(N^2)
// SC: O(N)
func waysToSplit(nums []int) int {
	// sums := make([]int, len(nums))
	// sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	n := len(nums) - 1

	l := 0
	r := 0
	answer := 0
	for i := 0; i < n-1; i++ {
		if l < i+1 {
			l = i + 1
		}

		for l < len(nums) && nums[l] < 2*nums[i] {
			l++
		}

		if r < l {
			r = l
		}

		for r < len(nums)-1 && nums[r]-nums[i] <= nums[n]-nums[r] {
			r++
		}

		answer = (answer + r - l) % 1000000007
	}

	return answer
}
