package main

import "fmt"

// TC: O(N + K)
// SC: O(K)
func countingSort(arr []int) {
	maxValue := findMaxValue(arr)

	var counts []int
	counts = make([]int, maxValue+1)

	for _, v := range arr {
		counts[v]++
	}

	var count, startingIndex int
	for i := 0; i < maxValue+1; i++ {
		count = counts[i]
		counts[i] = startingIndex
		startingIndex += count
	}

	var sortedArray []int
	sortedArray = make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		elem := arr[i]
		sortedArray[counts[elem]] = elem
		counts[elem]++
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = sortedArray[i]
	}
}

func countingSortWithShifting(arr []int) {
	shift := findMinValue(arr)
	maxValue := findMaxValue(arr) - shift

	var counts []int
	counts = make([]int, maxValue+1)

	for _, v := range arr {
		counts[v-shift]++
	}

	var count, startingIndex int
	for i := 0; i < maxValue+1; i++ {
		count = counts[i]
		counts[i] = startingIndex
		startingIndex += count
	}

	var sortedArray []int
	sortedArray = make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		elem := arr[i]
		sortedArray[counts[elem-shift]] = elem
		counts[elem-shift]++
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = sortedArray[i]
	}
}

func findMaxValue(arr []int) int {
	maxValue := arr[0]
	for _, v := range arr {
		if maxValue < v {
			maxValue = v
		}
	}

	return maxValue
}

func findMinValue(arr []int) int {
	minValue := arr[0]
	for _, v := range arr {
		if minValue > v {
			minValue = v
		}
	}

	return minValue
}

func main() {
	arr := []int{5, 3, 4, 4, 1, 9, -3, -2}
	countingSortWithShifting(arr)

	fmt.Println(arr)
}
