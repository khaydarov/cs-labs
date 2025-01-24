package main

type Num struct {
	value string
	index int
}

type NumList []Num

func toDigit(byte byte) int {
	return int(byte) - 48
}

func countSort(nums NumList, digitIndex int) {
	counts := make([]int, 10)
	for _, num := range nums {
		counts[toDigit(num.value[digitIndex])]++
	}

	startingIndex := 0
	for i := 0; i < 10; i++ {
		count := counts[i]
		counts[i] = startingIndex
		startingIndex += count
	}

	sorted := make([]Num, len(nums))
	for i := 0; i < len(nums); i++ {
		current := toDigit(nums[i].value[digitIndex])
		sorted[counts[current]] = nums[i]
		counts[current]++
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = sorted[i]
	}
}

func radixSort(nums NumList) {
	maxLength := len(nums[0].value)

	for maxLength > 0 {
		countSort(nums, maxLength-1)
		maxLength--
	}
}

func executeQuery(nums []string, k int) int {
	var numList NumList
	for i, num := range nums {
		numList = append(numList, Num{num, i})
	}

	radixSort(numList)

	return numList[k-1].index
}

func makeInput(nums []string, trim int) []string {
	var result []string
	for _, num := range nums {
		offset := len(num) - trim
		result = append(result, num[offset:])
	}

	return result
}

func SmallestTrimmedNumbers(nums []string, queries [][]int) []int {
	var result []int
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		trimmedInput := makeInput(nums, query[1])

		trimmedInputCopy := make([]string, len(trimmedInput))
		copy(trimmedInputCopy, trimmedInput)

		result = append(result, executeQuery(trimmedInputCopy, query[0]))
	}

	return result
}
