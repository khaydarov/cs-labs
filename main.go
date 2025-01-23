package main

import (
	"fmt"
)

type Num struct {
	num       int
	frequency int
}

type NumList []Num

func quickSelect(numList *NumList, left, right, k int) []int {
	pivot := partition(numList, left, right)
	if pivot == k {
		var result []int
		for i := 0; k > 0; i++ {
			result = append(result, (*numList)[i].num)
			k--
		}
		return result
	} else if pivot > k {
		return quickSelect(numList, left, pivot-1, k)
	} else {
		return quickSelect(numList, pivot+1, right, k)
	}
}

func partition(numList *NumList, left, right int) int {
	if left > right {
		return left
	}
	pivot := left + (right-left)/2
	(*numList)[right], (*numList)[pivot] = (*numList)[pivot], (*numList)[right]
	for i := left; i < right; i++ {
		if (*numList)[i].frequency > (*numList)[right].frequency {
			(*numList)[left], (*numList)[i] = (*numList)[i], (*numList)[left]
			left++
		}
	}

	(*numList)[right], (*numList)[left] = (*numList)[left], (*numList)[right]

	return left
}

func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}

	var numList NumList
	for k, v := range numMap {
		numList = append(numList, Num{k, v})
	}

	return quickSelect(&numList, 0, len(numList)-1, k)
}

func main() {
	nums := []int{1}
	r := topKFrequent(nums, 1)
	fmt.Println(r)
}
