package main

// TC: O(N)
// SC: O(1)
func findMaxConsecutiveOnes(nums []int) int {
	numZeros := 0

	left := 0
	right := 0

	answer := 0

	for right < len(nums) {
		if nums[right] == 0 {
			numZeros++
		}

		for numZeros > 1 {
			if nums[left] == 0 {
				numZeros--
			}
			left++
		}

		answer = max(answer, right-left+1)
		right++
	}

	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type Stream struct {
	left                   int
	right                  int
	lastZeroIndex          int
	longestConsecutiveOnes int
}

func (s *Stream) Add(x int) int {
	s.right++
	if x == 0 {
		for s.left < s.lastZeroIndex {
			s.left++
		}

		s.lastZeroIndex = s.right
	}

	s.longestConsecutiveOnes = max(s.longestConsecutiveOnes, s.right-s.left)

	return s.longestConsecutiveOnes
}
