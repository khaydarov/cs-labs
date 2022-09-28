package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 5, 5, 6, 6, 7, 7, 8, 9}
	r := findClosestElements(nums, 7, 9)
	fmt.Println(r)
}

func findClosestElements(arr []int, k int, x int) []int {
	low := 0
	high := len(arr) - k
	for low < high {
		mid := (low + high) / 2
		fmt.Println(low, high, mid)
		if x <= arr[mid] {
			high = mid
		} else if x >= arr[mid+k] {
			low = mid + 1
		} else {
			midList := abs(x - arr[mid])
			midkList := abs(x - arr[mid+k])
			if midList <= midkList {
				high = mid
			} else {
				low = mid + 1
			}
		}
	}

	return arr[low : low+k]
}

// TC: O(n-k)
// SC: O(1)
func findClosestElements1(arr []int, k int, x int) []int {
	low := 0
	high := len(arr) - 1

	for high-low >= k {
		if abs(arr[low]-x) <= abs(arr[high]-x) {
			high--
		} else {
			low++
		}
	}

	return arr[low : high+1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
