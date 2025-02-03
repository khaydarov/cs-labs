package main

// TC: O(N * 2^N)
// SC: O(2^N)
func grayCode1(n int) []int {
	seen := make(map[int]bool)
	seen[0] = true

	result := []int{0}
	var helper func(n int) bool
	helper = func(n int) bool {
		if len(result) >= 1<<n {
			return true
		}

		current := result[len(result)-1]
		for i := 0; i < n; i++ {
			next := current ^ (1 << i)
			if seen[next] {
				continue
			}

			result = append(result, next)
			seen[next] = true
			if helper(n) {
				return true
			}

			seen[next] = false
			result = result[:len(result)-1]
		}

		return false
	}

	helper(n)
	return result
}

// TC: O(2^N)
// SC: O(N)
func grayCode2(n int) []int {
	var result []int
	var helper func(result *[]int, n int)
	helper = func(result *[]int, n int) {
		if n == 0 {
			*result = append(*result, 0)
			return
		}

		helper(result, n-1)
		l := len(*result)
		mask := 1 << (n - 1)
		for i := l - 1; i >= 0; i-- {
			*result = append(*result, (*result)[i]|mask)
		}
	}

	helper(&result, n)
	return result
}

// TC: O(2^N)
// SC: O(1)
func grayCode3(n int) []int {
	result := []int{0}
	if n == 0 {
		return result
	}

	shift := 1
	for len(result) <= 1<<(n-1) {
		mask := 1 << (shift - 1)
		for i := len(result) - 1; i >= 0; i-- {
			result = append(result, result[i]|mask)
		}
		shift++
	}
	return result
}
