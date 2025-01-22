package insertion

// Complexity: best | average | worst
// TC: O(n) | O(n^2) | O(n^2)
// SC: O(1) | O(1) | O(1)
func insertSort(arr *[]int) {
	sortedRangeIndex := 1

	for sortedRangeIndex < len(*arr) {
		if (*arr)[sortedRangeIndex-1] > (*arr)[sortedRangeIndex] {
			insertionIndex := findInsertionIndex(arr, (*arr)[sortedRangeIndex])

			if insertionIndex == -1 {
				continue
			}

			insertValueAtIndex(arr, insertionIndex, sortedRangeIndex)
		}
		sortedRangeIndex++
	}
}

func findInsertionIndex(arr *[]int, element int) int {
	for i, v := range *arr {
		if v > element {
			return i
		}
	}

	return -1
}

func insertValueAtIndex(arr *[]int, insertAt int, insertFrom int) {
	if insertAt < 0 || insertFrom < 0 {
		return
	}

	insertingValue := (*arr)[insertFrom]

	for i := insertFrom; i > insertAt; i-- {
		(*arr)[i] = (*arr)[i-1]
	}

	(*arr)[insertAt] = insertingValue
}
