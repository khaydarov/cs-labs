package main

func hammingWeight1(num uint32) int {
	s := 0
	for num > 0 {
		if num&1 == 1 {
			s++
		}
		num >>= 1
	}

	return s
}

func hammingWeight2(num uint32) int {
	s := 0
	for num > 0 {
		s++
		num &= num - 1
	}

	return s
}
