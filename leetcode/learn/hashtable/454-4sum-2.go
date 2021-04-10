package main

import "fmt"

// TC: O(N^2)
// SC: O(N)
func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			s := a + b
			m[s]++
		}
	}

	result := 0
	for _, c := range C {
		for _, d := range D {
			s := c + d
			result += m[-1 * s]
		}
	}

	return result
}

func main() {
	a := []int{5}
	b := []int{-3}
	c := []int{-1}
	d := []int{-1}

	fmt.Println(fourSumCount(a, b, c, d))
}