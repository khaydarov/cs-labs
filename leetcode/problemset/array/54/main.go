package main

func SpiralCopy(inputMatrix [][]int) []int {
	rows := len(inputMatrix)
	cols := len(inputMatrix[0])

	visited := 101

	var result []int

	dirs := [][]int{{0, 1, cols - 1}, {1, 0, rows - 1}, {0, -1, cols - 1}, {-1, 0, rows - 1}}

	x, y, p := 0, 0, 1
	for p < rows*cols {
		for _, dir := range dirs {
			for i := 0; i < dir[2]; i++ {
				if inputMatrix[x][y] == visited {
					continue
				}
				result = append(result, inputMatrix[x][y])
				inputMatrix[x][y] = visited
				x += dir[0]
				y += dir[1]
				p++
			}
			dir[2] -= 2
		}
		x++
		y++
	}

	if p == rows*cols {
		result = append(result, inputMatrix[x][y])
	}
	
	return result
}
