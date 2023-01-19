package main

import "math"

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(c int) {
	q.data = append(q.data, c)
}

func (q *Queue) Dequeue() {
	q.data = q.data[1:]
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue) Front() int {
	return q.data[0]
}

func (q *Queue) Length() int {
	return len(q.data)
}

// TC: O(N^h)
// SC: O(sqrt(N)^h); where h is the height of tree
func numSquares(n int) int {
	q := Queue{}
	q.Enqueue(n)

	result := 0
	for !q.Empty() {
		l := q.Length()

		identicals := make(map[int]bool)
		for i := 0; i < l; i++ {
			current := q.Front()

			if current == 0 {
				return result
			}

			for j := 1; j*j <= current; j++ {
				if _, ok := identicals[current-j*j]; !ok {
					q.Enqueue(current - j*j)
					identicals[current-j*j] = true
				}

			}

			q.Dequeue()
		}

		result++
	}

	return -1
}

// TC: O(N * sqrt(N))
// SC: O(N)
func numSquaresDP(n int) int {
	// [0, 1, 2, 3, 4, 5, 6, 7, 8, ..., n]
	// [0, 1, 2, 3, 1, 2, 3, ...]

	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			if dp[i] > dp[i-j*j]+1 {
				dp[i] = dp[i-j*j] + 1
			}
		}
	}

	return dp[n]
}

func numSquaresMath(n int) int {
	for n%4 == 0 {
		n /= 4
	}

	if n%8 == 7 {
		return 4
	}

	if isSquare(n) {
		return 1
	}

	for i := 1; i*i <= n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	return 3
}

func isSquare(n int) bool {
	r := int(math.Sqrt(float64(n)))

	return r*r == n
}
