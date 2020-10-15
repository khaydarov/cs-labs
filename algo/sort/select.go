package main

func main() {
	arr := []int{3, 7, 4, 4, 6, 5, 8}
	selectSort(&arr)
}

// Complexity: best | average | worst
// TC: O(n) | O(n^2) | O(n^2)
// SC: O(1) | O(1) | O(1)
func selectSort(arr *[]int) {
	sortIndex := 0

	for sortIndex < len(*arr) {
		index := findSmallestFromIndex(arr, sortIndex)
		swap(arr, sortIndex, index)

		sortIndex++
	}
}

func findSmallestFromIndex(arr *[]int, index int) int {
	min := (*arr)[index]

	for i := index + 1; i < len(*arr); i++ {
		if min > (*arr)[i] {
			min = (*arr)[i]
			index = i
		}
	}
	return index
}

func swap(arr *[]int, a int, b int) {
	(*arr)[a], (*arr)[b] = (*arr)[b], (*arr)[a]
}