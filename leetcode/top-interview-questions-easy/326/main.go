package main

// Naive
// TC: O(Log3 N)
// SC: O(1)
func isPowerOfThree1(n int) bool {
	if n <= 0 {
		return false
	}

	k := 1
	x := 0
	for k < n {
		k *= 3
		x++
	}

	return k == n
}

// Math: Max integer value is 2^31 - 1. So the max value that is power of three is 3^19: 3^log3 2^31-1 ~ 3^19,...
// TC: O(1)
// SC: O(1)
func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}
