package main

import "sort"

// TC: O(N * M)
// SC: O(1)
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)

	s := 0
	for _, v := range arr1 {
		found := false
		for i := len(arr2) - 1; i >= 0; i-- {
			if abs(arr2[i]-v) <= d {
				found = true
				break
			}
		}

		if !found {
			s++
		}
	}

	return s
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
