package main

import "fmt"

func main() {
	arr := []int{10,9,1,8}
	mergeSort(&arr)

	fmt.Println("result", arr)
}

// Complexity: best | average | worst
// TC: O(n log n) | O(n log n) | O(n log n)
// SC: O(n) | O(n) | O(n)
func mergeSort(arr *[]int) {
	length := len(*arr)

	if length <= 1 {
		return
	}

	leftSize := length / 2

	left := make([]int, leftSize)
	for i := 0; i < leftSize; i++ {
		left[i] = (*arr)[i]
	}

	right := make([]int, length - leftSize)
	for i := 0; i < length - leftSize; i++ {
		right[i] = (*arr)[length - i - 1]
	}

	mergeSort(&left)
	mergeSort(&right)
	merge(arr, left, right)
}

func merge(arr *[]int, left []int, right []int) {
	leftIndex := 0
	RightIndex := 0
	targetIndex := 0
	remaining := len(left) + len(right)

	fmt.Println(left, right)
	for remaining > 0 {
		if leftIndex >= len(left) {
			(*arr)[targetIndex] = right[RightIndex]
			RightIndex++
		} else if RightIndex >= len(right) {
			(*arr)[targetIndex] = left[leftIndex]
			leftIndex++
		} else if left[leftIndex] < right[RightIndex] {
			(*arr)[targetIndex] = left[leftIndex]
			leftIndex++
		} else {
			(*arr)[targetIndex] = right[RightIndex]
			RightIndex++
		}

		targetIndex++
		remaining--
	}
}
