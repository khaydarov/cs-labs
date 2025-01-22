package heap

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
func HeapSort(arr *[]int) {
	n := len(*arr)

	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, n)
	}

	for i := n - 1; i >= 0; i-- {
		(*arr)[0], (*arr)[i] = (*arr)[i], (*arr)[0]
		maxHeapify(arr, 0, i)
	}
}
