package main

import "fmt"

func main() {
	k := fib(7)

	fmt.Println(k)
}

func fib(n int) int {
	base := [][]int{{0, 1}, {1, 1}}
	matrix := matrixPow(base, n)

	return matrix[1][0]
}

func matrixPow(a [][]int, n int) [][]int {
	if n == 0 {
		return [][]int{{0, 1}}
	}

	if n == 1 {
		return a
	}

	result := [][]int{{1, 0}, {0, 1}}
	for n > 0 {
		if n % 2 == 1 {
			result = matrixMultiply(result, a)
		}
		a = matrixMultiply(a, a)
		n /= 2
	}

	return result
}

func matrixMultiply(a, b [][]int) [][]int {
	c := make([][]int, 2)
	for i := 0; i < 2; i++ {
		c[i] = make([]int, 2)
	}

	c[0][0] = a[0][0] * b[0][0] + a[0][1] * b[1][0]
	c[0][1] = a[0][0] * b[0][1] + a[0][1] * b[1][1]
	c[1][0] = a[1][0] * b[0][0] + a[1][1] * b[1][0]
	c[1][1] = a[1][0] * b[0][1] + a[1][1] * b[1][1]

	return c
}