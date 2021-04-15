package main

import (
	"fmt"
)

// TC: O(2^N)
// SC: O(2^N)
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	result := countAndSay(n - 1)

	return countHelper(result)
}

func countHelper(s string) string {
	b := s[0]
	count := 1
	result := ""

	var countByte byte
	for i := 1; i < len(s); i++ {
		if s[i] == b {
			count++
		} else {
			countByte = byte(count + '0')
			result += string(countByte) + string(b)

			b = s[i]
			count = 1
		}
	}

	countByte = byte(count + '0')
	result += string(countByte) + string(b)

	return result
}

func main() {

	fmt.Println(countAndSay(7))
}