package main

func generate(numRows int) [][]int {
	var result [][]int
	for i := 0; i < numRows; i++ {
		var row []int
		row = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = result[i-1][j-1] + result[i-1][j]
			}
		}

		result = append(result, row)
	}

	return result
}
