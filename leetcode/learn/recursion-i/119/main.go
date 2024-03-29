package main

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 0; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			result[j] = result[j] + result[j-1]
		}
	}
	result[rowIndex] = 1

	return result
}
