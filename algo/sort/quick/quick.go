package quick

import (
	"math/rand"
)

// Complexity: best | average | worst
// TC: O(n log n) | O(n log n) | O(n^2)
// SC: O(n) | O(n) | O(n)
func quickSort(arr *[]int) {
	if len(*arr) < 2 {
		return
	}

	left, right := 0, len(*arr)-1
	pivot := rand.Int() % len(*arr)

	(*arr)[pivot], (*arr)[right] = (*arr)[right], (*arr)[pivot]

	for i, _ := range *arr {
		if (*arr)[i] < (*arr)[right] {
			(*arr)[left], (*arr)[i] = (*arr)[i], (*arr)[left]
			left++
		}
	}

	(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]

	l := (*arr)[:left]
	r := (*arr)[left+1:]
	quickSort(&l)
	quickSort(&r)
}
