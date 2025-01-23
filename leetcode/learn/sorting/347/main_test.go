package main

import "testing"

func TestSolution(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	r := topKFrequent(nums, 2)
	if r[0] != 1 && r[1] != 2 {
		t.Fatal("Test failed")
	}
}
