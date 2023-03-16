package main

import "math"

// Return the size of bit array(m) to used using
// Following formula
// m = -(n * lg(p)) / (lg(2)^2)
// Where
// n: int - number of items expected to be stored in filter
// p: float - False Positive probability in decimal
func calculateBFSize(n int, p float64) int {
	m := -(float64(n) * math.Log(p)) / math.Pow(math.Log(2), 2)

	return int(m)
}

// Return the hash function(k) to be used using
// Following formula
// k = (m / n) * lg(2)
// m : int - size of bit array
// n : int - number of items expected to be stored in filter
func calculateHashCount(n, m int) int {
	k := float64(m/n) * math.Log(2)

	return int(k)
}
