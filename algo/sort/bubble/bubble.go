package bubble

// Complexity: best | average | worst
// TC: O(n) | O(n^2) | O(n^2)
// SC: O(1) | O(1) | O(1)
func bubbleSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[i] > (*arr)[j] {
				(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			}
		}
	}
}
