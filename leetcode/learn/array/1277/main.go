package main

import "fmt"

func main() {
	//matrix := [][]int{ {1, 0, 1}, {1, 1, 0}, {1, 1, 0}}
	matrix := [][]int{{0, 1}, {1, 1}}
	r := countSquares(matrix)
	fmt.Println(r)
}

func countSquares(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])

	// Using dynamic programming
	// Consider that each matrix element is the right-bottom edge of the square
	// For the squares of size more that one we check matrix[i - 1][j], matrix[i][j - 1], matrix[i - 1][j - 1]
	// And get minimum of these values: that will mean there can not be a square
	// Ex.
	// 0 1
	// 1 1
	// In the example above we can see that "right-bottom" element is one, and it can be a square of size 1x1, but not more
	// So, the solution is creating the additive matrix with sums
	// dp[i][j] = min(dp[i][j - 1], dp[i - 1][j], dp[i - 1][j - 1]) + 1
	// and the result is sum of all additive elements
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 && i > 0 && j > 0 {
				matrix[i][j] = min(matrix[i-1][j], min(matrix[i][j-1], matrix[i-1][j-1])) + 1
			}
			result += matrix[i][j]
		}
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
