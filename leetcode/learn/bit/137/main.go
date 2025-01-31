package main

// TC: O(N) - N is the number of elements in the input array
// SC: O(1)
func singleNumber(nums []int) int {
	loner := 0
	for i := 0; i < 32; i++ {
		bitsSum := 0
		for _, num := range nums {
			bitsSum += (num >> i) & 1
		}

		lonerBit := bitsSum % 3
		loner |= lonerBit << i
	}

	if loner >= (1 << 31) {
		loner -= (1 << 32)
	}

	return loner
}

// Repeat Approach 4
// TC: O(N) - N is the number of elements in the input array
// SC: O(1)
func singleNumber1(nums []int) int {
	seenOnce, seenTwice := 0, 0
	for _, num := range nums {
		seenOnce = (seenOnce ^ num) & (^seenTwice)
		seenTwice = (seenTwice ^ num) & (^seenOnce)
	}

	return seenOnce
}
