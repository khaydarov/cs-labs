package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	total := len(nums1) + len(nums2)
	half := (total + 1) / 2

	var Aleft, Aright float64
	var Bleft, Bright float64

	low, high := 0, len(nums1)-1
	for {
		partitionX := (low + high) / 2
		partitionY := half - partitionX - 2

		if partitionX >= 0 {
			Aleft = float64(nums1[partitionX])
		} else {
			Aleft = math.Inf(-1)
		}

		if (partitionX + 1) < len(nums1) {
			Aright = float64(nums1[partitionX+1])
		} else {
			Aright = math.Inf(1)
		}

		if partitionY >= 0 {
			Bleft = float64(nums2[partitionY])
		} else {
			Bleft = math.Inf(-1)
		}

		if (partitionY + 1) < len(nums2) {
			Bright = float64(nums2[partitionY+1])
		} else {
			Bright = math.Inf(1)
		}

		fmt.Println(partitionX, partitionY, Aleft, Aright, Bleft, Bright)

		// partition is correct
		if Aleft <= Bright && Bleft <= Aright {
			// odd
			if total%2 == 1 {
				return max(Aleft, Bleft)
			}
			// even
			return (max(Aleft, Bleft) + min(Aright, Bright)) / 2
		} else if Aleft > Bright {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	return 1.0
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	r := findMedianSortedArrays(
		[]int{},
		[]int{1},
	)

	fmt.Println(r)
}
