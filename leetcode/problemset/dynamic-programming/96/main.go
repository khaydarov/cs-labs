package main

// TC: O(N^2)
// SC: O(N)
func numTrees1(n int) int {
	g := make([]int, n+1)
	g[0] = 1
	g[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}

	return g[n]
}

// TC: O(N)
// SC: O(1)
func numTrees(n int) int {
	var c int
	c = 1

	for i := 0; i < n; i++ {
		c = 2 * c * (2*i + 1) / (i + 2)
	}

	return c
}
