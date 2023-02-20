package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	var result []int
	sort.Ints(nums1)
	sort.Ints(nums2)

	i := 0
	j := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1, 4, 4}
	nums2 := []int{2, 2, 4}

	r := intersect(nums1, nums2)

	fmt.Println(r)
}
