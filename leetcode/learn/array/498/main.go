package main

import "fmt"

// TC: O(N + M)
// SC: O(1) or O(N + M)
func findDiagonalOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	var result []int

	n := len(matrix)
	m := len(matrix[0])

	i := 0
	j := 0
	for len(result) < m * n {
		if i == n - 1 && j == m - 1 {
			break
		}

		for true {
			result = append(result, matrix[i][j])
			if i == 0 || j == m - 1 {
				break
			}
			i--
			j++
		}

		if i == 0 && j < m - 1 {
			j++
		} else {
			i++
		}

		for true {
			result = append(result, matrix[i][j])
			if j == 0 || i == n - 1 {
				break
			}
			j--
			i++
		}

		if j == 0 && i < n - 1 {
			i++
		} else {
			j++
		}
	}

	return result
}

// TC: O(N + M)
// SC: O(N + M)
func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	n := len(matrix)
	m := len(matrix[0])

	diagonals := make(map[int][]int, n * m + 1)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			diagonals[i + j] = append(diagonals[i + j], matrix[i][j])
		}
	}

	var result []int
	for i := 0; i < n * m; i++ {
		diagonal := diagonals[i]

		if i % 2 == 1 {
			result = append(result, diagonal...)
		} else {
			result = append(result, reverse(diagonal)...)
		}
	}

	return result
}

func reverse(arr []int) []int {
	n := len(arr)
	for i := 0; i < n / 2; i++ {
		arr[i], arr[n - i - 1] = arr[n - i - 1], arr[i]
	}

	return arr
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{6, 7, 8},
		{11, 12, 13},
		{16, 17, 18},
	}
	r := findDiagonalOrder(matrix)
	fmt.Println(r)
}