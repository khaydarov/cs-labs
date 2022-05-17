package main

// TC: O(N)
// SC: O(1)
func maxSum(arr []int, k int) int {
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	answer := windowSum
	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i - k]

		if answer < windowSum {
			answer = windowSum
		}
	}

	return answer
}


func main() {
	maxSum([]int{5, 2, -1, 0, 3}, 3)
}

