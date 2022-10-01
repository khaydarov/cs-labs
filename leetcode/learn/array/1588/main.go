package main

// TC: O(N^2)
// SC: O(N)
func sumOddLengthSubarrays(arr []int) int {
	prefix := make([]int, len(arr)+1)
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}

	var answer int
	for length := 1; length <= len(arr); length += 2 {
		for i := 0; i < len(arr); i++ {
			if i+length > len(arr) {
				break
			}

			answer += prefix[i+length] - prefix[i]
		}
	}

	return answer
}
