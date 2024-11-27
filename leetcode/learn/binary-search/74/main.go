package main

// TC: O(log (N * M))
// SC: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row := searchRow(matrix, 0, len(matrix)-1, target)
	if row == -1 {
		return false
	}

	if matrix[row][0] == target {
		return true
	}

	col := searchCol(matrix, row, 0, len(matrix[0])-1, target)
	if col == -1 {
		return false
	}

	return matrix[row][col] == target
}

func searchRow(matrix [][]int, low, high, target int) int {
	for low <= high {
		mid := low + (high-low)/2
		if matrix[mid][0] == target {
			return mid
		}

		if matrix[mid][0] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return high
}

func searchCol(matrix [][]int, row, low, high, target int) int {
	for low <= high {
		mid := low + (high-low)/2
		if matrix[row][mid] == target {
			return mid
		}

		if matrix[row][mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// TC: O(log (N*M))
// SC: O(1)
func searchMatrix1(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	if rows == 0 || cols == 0 {
		return false
	}

	low := 0
	high := rows*cols - 1
	for low <= high {
		pivot := low + (high-low)/2
		pivotElement := matrix[pivot/cols][pivot%cols]
		if pivotElement == target {
			return true
		} else if pivotElement < target {
			low = pivot + 1
		} else {
			high = pivot - 1
		}
	}

	return false
}
