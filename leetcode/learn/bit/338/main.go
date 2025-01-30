package main

// TC: O(nlogn)
// SC: O(1)
func countBits1(n int) []int {
	helper := func(k int) int {
		count := 0
		for k != 0 {
			k &= (k - 1)
			count++
		}
		return count
	}

	var result []int
	for i := 0; i <= n; i++ {
		result = append(result, helper(i))
	}

	return result
}

// TC: O(n)
// SC: O(1)
func countBits2(n int) []int {
	result := make([]int, n+1)
	var x, b int
	b = 1

	for b <= n {
		for x < b && x+b <= n {
			result[x+b] = result[x] + 1
			x++
		}

		x = 0
		b <<= 1
	}

	return result
}

func countBits3(n int) []int {
	result := make([]int, n+1)
	result[0] = 0
	result[1] = 1

	for i := 2; i <= n; i++ {
		result[i] = result[i/2] + result[i%2]
	}

	return result
}
