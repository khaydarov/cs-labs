package main

import "fmt"

// TC: O(N), where N is the length of string
// SC: O(1)
func myAtoi(s string) int {
	s = trimLeadingSpaces(s)

	if len(s) == 0 {
		return 0
	}

	if !isDigit(s[0]) && !isSign(s[0]) {
		return 0
	}

	if isSign(s[0]) {
		signCount := 0
		for signCount < len(s) && isSign(s[signCount]) {
			signCount++
		}

		if signCount > 1 {
			return 0
		}
	}

	i := 1
	for i < len(s) && isDigit(s[i]) {
		i++
	}

	s = s[:i]

	min := -2147483648
	max := 2147483647

	sign := 1
	digitsStartPosition := 0
	if s[0] == '-' {
		sign = -1
		digitsStartPosition++
	} else if s[0] == '+' {
		sign = 1
		digitsStartPosition++
	}

	sum := 0
	for i := digitsStartPosition; i < len(s); i++ {
		digit := int(s[i] - '0')

		if sum > max / 10 || (sum == max / 10 && digit > 7) {
			return max
		}

		if sum < min / 10 || (sum == min / 10 && digit > 8) {
			return min
		}

		if sign == 1 {
			sum = sum * 10 + digit
		} else if sign == -1 {
			sum = sum * 10 - digit
		}
	}

	return sum
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSign(b byte) bool {
	return b == '+' || b == '-'
}

func trimLeadingSpaces(s string) string {
	if len(s) == 0 {
		return s
	}

	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	return s[i:]
}

func main() {
	//s := "    -42"
	//s := "    -4a2"
	//s := "    4a2"
	//s := "    --4a2"
	//s := "    +-4a2"
	//s := "    +-4a2"
	//s := "    +42"
	//s := "4193 with words"
	//s := "words and 987"
	//s := "-91283472332"
	//s := "91283472332"
	//s := ""
	//s := "     "
	//s := "- 3-3"
	//s := "-2147483647"
	s := "2147483648"
	fmt.Println(myAtoi(s))
}