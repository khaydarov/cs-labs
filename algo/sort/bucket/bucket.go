package bucket

import (
	"sort"
)

func findMax(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func findMin(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func BucketSort(nums []int) {
	numBuckets := len(nums) / 2
	buckets := make([][]int, numBuckets+1)

	shift := findMin(nums)
	maxValue := findMax(nums) - shift
	bucketSize := max(maxValue/numBuckets, 1)

	for _, num := range nums {
		index := (num - shift) / bucketSize
		buckets[index] = append(buckets[index], num)
	}

	for i := 0; i <= numBuckets; i++ {
		sort.Slice(buckets[i], func(a, b int) bool {
			return buckets[i][a] < buckets[i][b]
		})
	}

	sortedNums := make([]int, 0, len(nums))
	for _, bucket := range buckets {
		sortedNums = append(sortedNums, bucket...)
	}

	copy(nums, sortedNums)
}
