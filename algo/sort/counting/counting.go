package counting

// K is the maximum value in the array
// TC: O(N + K)
// SC: O(K)
func CountingSort(arr []int) {
	maxValue := findMaxValue(arr)

	counts := make([]int, maxValue+1)
	for _, v := range arr {
		counts[v]++
	}

	var count, startingIndex int
	for i := 0; i < maxValue+1; i++ {
		count = counts[i]
		counts[i] = startingIndex
		startingIndex += count
	}

	sortedArray := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		elem := arr[i]
		sortedArray[counts[elem]] = elem
		counts[elem]++
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = sortedArray[i]
	}
}

func CountingSortWithShifting(arr []int) {
	shift := findMinValue(arr)
	maxValue := findMaxValue(arr) - shift

	counts := make([]int, maxValue+1)
	for _, v := range arr {
		counts[v-shift]++
	}

	var count, startingIndex int
	for i := 0; i < maxValue+1; i++ {
		count = counts[i]
		counts[i] = startingIndex
		startingIndex += count
	}

	sortedArray := make([]int, len(arr))
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
