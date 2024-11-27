package main

import "fmt"

func maxHeapify(arr *[]int, i int, heapSize int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < heapSize && (*arr)[left] > (*arr)[largest] {
		largest = left
	}

	if right < heapSize && (*arr)[right] > (*arr)[largest] {
		largest = right
	}

	if largest != i {
		(*arr)[i], (*arr)[largest] = (*arr)[largest], (*arr)[i]
		maxHeapify(arr, largest, heapSize)
	}
}

// Complexity: best | average | worst
// TC: O(NlogN) | O(NlogN) | O(NlogN)
// SC: O(1) | O(1) | O(1)
func heapSort(arr *[]int) {
	n := len(*arr)

	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, n)
	}

	for i := n - 1; i >= 0; i-- {
		(*arr)[0], (*arr)[i] = (*arr)[i], (*arr)[0]
		maxHeapify(arr, 0, i)
	}
}

func main() {
	arr := []int{5, 4, 3, 1, 2, 9}
	heapSort(&arr)
	fmt.Println(arr)
}
