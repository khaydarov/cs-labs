package main

// TC: O(N) - N is the length of nums
// SC: O(1) - only constant space is used
func singleNumber(nums []int) []int {
	bitMask := 0
	for _, num := range nums {
		bitMask ^= num
	}

	diff := bitMask & (-bitMask)
	x := 0
	for _, num := range nums {
		if num&diff != 0 {
			x ^= num
		}
	}

	return []int{x, bitMask ^ x}
}
