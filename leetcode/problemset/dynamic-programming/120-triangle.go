package main

import "fmt"

// TC: O(N^2)
// SC: O(N)
func minimumTotal(triangle [][]int) int {
	paths := make([]int, len(triangle))
	paths[0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0 ; j-- {
			c := triangle[i][j]
			if j == 0 {
				paths[j] += c
			} else if j == len(triangle[i]) - 1{
				paths[j] = paths[j - 1] + c
			} else {
				paths[j] = min(paths[j] + c, paths[j - 1] + c)
			}
		}
	}

	min := paths[0]
	for i := 1; i < len(paths); i++ {
		if min > paths[i] {
			min = paths[i]
		}
	}

	fmt.Println(min, paths)
	return min
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 2, 1},
		{4, 1, 8, 2},
	}

	minimumTotal(triangle)
}