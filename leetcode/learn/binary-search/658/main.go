package main

func findClosestElements(arr []int, k int, x int) []int {
	left := 0
	right := len(arr) - k

	for left < right {
		pivot := left + (right-left)/2

		if arr[pivot] == arr[pivot+k] && x > arr[pivot] {
			left = pivot + 1
		} else if abs(x-arr[pivot]) > abs(arr[pivot+k]-x) {
			left = pivot + 1
		} else {
			right = pivot
		}
	}

	return arr[left : left+k]
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}
