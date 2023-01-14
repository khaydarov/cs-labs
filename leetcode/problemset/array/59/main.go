package main

func generateMatrix(n int) [][]int {
	var matrix [][]int
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}

	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	x, y, p := 0, 0, 1

	offset := 1
	for p < n*n {
		for _, dir := range dirs {
			for i := 0; i < n-offset; i++ {
				matrix[x][y] = p
				p++
				x += dir[0]
				y += dir[1]
			}
		}

		x++
		y++
		offset += 2
	}

	if offset == n {
		matrix[x][y] = p
	}

	return matrix
}
