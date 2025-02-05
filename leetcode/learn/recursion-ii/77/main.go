package main

func combine(n int, k int) [][]int {
	var result [][]int

	current := make([]int, k)
	var backtrack func(startPos, startValue int)
	backtrack = func(startPos, startValue int) {
		if startPos == k {
			combo := make([]int, k)
			copy(combo, current)
			result = append(result, combo)

			return
		}

		maxStart := n - (k - startPos) + 1
		for i := startValue; i <= maxStart; i++ {
			current[startPos] = i
			backtrack(startPos+1, i+1)
		}
	}
	backtrack(0, 1)

	return result
}
