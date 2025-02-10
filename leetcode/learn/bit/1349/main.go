package main

func maxStudents(seats [][]byte) int {
	n := len(seats)
	m := len(seats[0])

	var dp func(position, previousRow, currentRow int) int
	dp = func(position, previousRow, currentRow int) int {
		if position == m*n {
			return 0
		}

		i := position / m
		j := position % m

		if j == 0 {
			previousRow = currentRow
			currentRow = 0
		}

		ans := 0
		ans = max(ans, dp(position+1, previousRow, currentRow))

		if seats[i][j] == '.' {
			left, topLeft, right, topRight := true, true, true, true

			if j > 0 {
				left = (currentRow & (1 << (j - 1))) == 0
				if i > 0 {
					topLeft = (previousRow & (1 << (j - 1))) == 0
				}
			}

			if j < m-1 {
				right = (currentRow & (1 << (j + 1))) == 0
				if i > 0 {
					topRight = (previousRow & (1 << (j + 1))) == 0
				}
			}

			if left && topLeft && right && topRight {
				ans = max(ans, dp(position+1, previousRow, currentRow|(1<<j))+1)
			}
		}

		return ans
	}

	return dp(0, 0, 0)
}
