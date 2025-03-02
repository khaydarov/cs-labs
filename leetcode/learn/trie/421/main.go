package main

import (
	"slices"
	"strconv"

	"github.com/emirpasic/gods/sets/hashset"
)

func findMaximumXOR1(nums []int) int {
	maxValue := slices.Max(nums)

	l := len(strconv.FormatInt(int64(maxValue), 2))
	set := hashset.New()

	maxXor := 0
	for i := l - 1; i >= 0; i-- {
		maxXor <<= 1
		current := maxXor | 1

		for _, num := range nums {
			set.Add(num >> i)
		}

		for _, prefix := range set.Values() {
			if set.Contains(prefix.(int) ^ current) {
				maxXor = current
			}
		}

		set.Clear()
	}

	return maxXor
}

type Node struct {
	Children [2]*Node
}

// Time complexity, O(n * l), where n is the number of elements in the array and l is the number of bits in the maximum value.
// Space complexity, O(n * l), where n is the number of elements in the array and l is the number of bits in the maximum value.
func findMaximumXOR2(nums []int) int {
	maxValue := slices.Max(nums)

	l := len(strconv.FormatInt(int64(maxValue), 2))
	root := &Node{}

	maxXor := 0
	for _, num := range nums {
		node := root
		xorNode := root
		currentXor := 0

		for i := l - 1; i >= 0; i-- {
			bit := num >> i & 1
			if node.Children[bit] == nil {
				node.Children[bit] = &Node{}
			}
			node = node.Children[bit]

			oppositeBit := 1 - bit
			if xorNode.Children[oppositeBit] != nil {
				xorNode = xorNode.Children[oppositeBit]
				currentXor |= 1 << i
			} else {
				xorNode = xorNode.Children[bit]
			}
		}

		if currentXor > maxXor {
			maxXor = currentXor
		}
	}

	return maxXor
}
