package main

// TC: O(right - left)
// SC: O(1)
func rangeBitwiseAnd1(left int, right int) int {
	result := left
	for i := left + 1; i <= right; i++ {
		result &= i
		if result == 0 {
			return 0
		}
	}

	return result
}

// TC: O(1)
// SC: O(1)
func rangeBitwiseAnd2(left int, right int) int {
	i := 0
	for left != right {
		left >>= 1
		right >>= 1
		i++
	}

	return right << i
}

// TC: O(1)
// SC: O(1)
func rangeBitwiseAnd3(left int, right int) int {
	for left < right {
		right = right & (right - 1)
	}

	return right
}
