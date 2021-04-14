package main

func rotate(matrix [][]int)  {
	n := len(matrix)
	for i := 0; i < n / 2; i++ {
		for j := i; j < n - i - 1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n - j - 1][i]
			matrix[n - j - 1][i] = matrix[n - i - 1][n - j - 1]
			matrix[n - i - 1][n - j - 1] = matrix[j][n - i - 1]
			matrix[j][n - i - 1] = tmp
		}
	}
}

func main() {
	//mx := [][]int{
	//	{1, 2, 3, 4, 5},
	//	{6, 7, 8, 9, 10},
	//	{11, 12, 13, 14, 15},
	//	{16, 17, 18, 19, 20},
	//	{21, 22, 23, 24, 25},
	//}

	mx := [][]int{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36},
	}
	rotate(mx)
}