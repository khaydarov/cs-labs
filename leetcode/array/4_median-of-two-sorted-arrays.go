package main

import (
	"fmt"
)

func main() {
	nums1 := []int{100000}
	nums2 := []int{100001}
	r := findMedianSortedArrays(nums1, nums2)
	fmt.Println(r)
}

const MAX = 1000001
const MIN = -1000001

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Guarantee that nums1 is smaller than nums2
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	// Store the length of arrays as X and Y
	x := len(nums1)
	y := len(nums2)

	low := 0
	high := x
	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x + y + 1) / 2 - partitionX

		// We get X[l] and X[r] values
		maxLeftX := MIN
		if partitionX != 0 {
			maxLeftX = nums1[partitionX - 1]
		}

		minRightX := MAX
		if partitionX != x {
			minRightX = nums1[partitionX]
		}

		// We get Y[l] and R[r] values
		maxLeftY := MIN
		if partitionY != 0 {
			maxLeftY = nums2[partitionY - 1]
		}

		minRightY := MAX
		if partitionY != y {
			minRightY = nums2[partitionY]
		}

		//fmt.Println(low, high)
		//fmt.Println(maxLeftX, maxLeftY, minRightX, minRightY)
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x + y) % 2 == 0 {
				// If X + Y is even than ans: avg(max(maxLeftX, maxLeftY), min(minRightX, minRightY))
				max := max(maxLeftX, maxLeftY)
				min := min(minRightX, minRightY)
				return float64(max + min) / 2
			} else {
				// if X + Y is odd than ans: max(maxLeftX, maxLeftY)
				max := max(maxLeftX, maxLeftY)
				return float64(max)
			}
		} else if maxLeftX >= minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	return -1.0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}