package main

func main() {

	arr := []int{0, 0, 0, 0}
	canMakeArithmeticProgression(arr)
}

func canMakeArithmeticProgression(arr []int) bool {
	n := len(arr)
	quickSort(&arr, 0, n - 1)

	if n == 2 {
		return true
	}

	d := arr[1] - arr[0]
	for i := 2; i < n; i++ {
		if arr[i] != arr[i - 1] + d {
			return false
		}
	}

	return true
}

func quickSort(arr *[]int, left int, right int) {
	if left >= right {
		return
	}

	pivot := (*arr)[left + (right-left)/2]
	index := partition(arr, left, right, pivot)

	quickSort(arr, left, index - 1)
	quickSort(arr, index, right)
}

func partition(arr *[]int, left, right, pivot int) int {
	for left <= right {
		for (*arr)[left] < pivot {
			left++
		}

		for (*arr)[right] > pivot {
			right--
		}

		if left <= right {
			(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
			left++
			right--
		}
	}

	return left
}