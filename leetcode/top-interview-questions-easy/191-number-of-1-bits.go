package main

import "fmt"

func hammingWeight(num uint32) int {
	s := 0
	for num > 0 {
		if num & 1 == 1 {
			s++
		}
		num >>= 1
	}

	return s
}

func main() {
	n := 1
	fmt.Println(n)
	fmt.Println(hammingWeight(uint32(n)))
}