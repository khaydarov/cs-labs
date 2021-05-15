package main

import "fmt"

// TC: O(N)
// SC: O(1)
func isNumber(s string) bool {
	if len(s) == 1 && !isCharDigit(s[0]) {
		return false
	}

	if containsLetters(s) {
		return false
	}

	if isInteger(s) || isDecimal(s) {
		return true
	}

	i := 0
	for i < len(s) && !isCharEpsilon(s[i]) {
		i++
	}

	if i == len(s) - 1 {
		return false
	}

	return (isInteger(s[:i]) || isDecimal(s[:i])) && isInteger(s[i+1:])
}

// TC: O(N)
// SC: O(1)
func isInteger(s string) bool {
	if s == "" {
		return false
	}

	if len(s) == 1 && !isCharDigit(s[0]) {
		return false
	}

	for i := 0; i < len(s); i++ {
		char := s[i]
		if isCharDigit(char) {
			continue
		}

		if isCharSign(char) && i > 0 {
			return false
		}

		if isCharEpsilon(char) || isCharDot(char) {
			return false
		}
	}

	return true
}

// TC: O(N)
// SC: O(1)
func isDecimal(s string) bool {
	if s == "" {
		return false
	}

	if len(s) == 1 && !isCharDigit(s[0]) {
		return false
	}

	hasDot := false
	for i := 0; i < len(s); i++ {
		char := s[i]

		if isCharDigit(char) {
			continue
		}

		if isCharSign(char) && i > 0 {
			return false
		}

		if isCharEpsilon(char) {
			return false
		}

		if isCharDot(char) {
			if hasDot {
				return false
			} else {
				hasDot = true
			}

			if i == 0 && !isCharDigit(s[i + 1]) {
				return false
			}

			if i == len(s) - 1 && !isCharDigit(s[i - 1]) {
				return false
			}
		}
	}

	return true
}

// TC: O(N), where N is length of string
// SC: O(1)
func containsLetters(s string) bool {
	for i := 0; i < len(s); i++ {
		v := s[i]
		if isCharLetter(v) && !isCharEpsilon(v) {
			return true
		}
	}

	return false
}

// TC: O(1)
// SC: O(1)
func isCharSign(b byte) bool {
	return b == '+' || b == '-'
}

// TC: O(1)
// SC: O(1)
func isCharDot(b byte) bool {
	return b == '.'
}

// TC: O(1)
// SC: O(1)
func isCharEpsilon(b byte) bool {
	return b == 'e' || b == 'E'
}

// TC: O(1)
// SC: O(1)
func isCharLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// TC: O(1)
// SC: O(1)
func isCharDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {
	tests := []struct{
		s string
		r bool
	}{
		{
			".1.",
			false,
		},
		{
			".e1",
			false,
		},
		{
			"0",
			true,
		},
		{
			"0089",
			true,
		},
		{
			"-0.1",
			true,
		},
		{
			"3.14",
			true,
		},
		{
			"-3.14",
			true,
		},
		{
			"4.",
			true,
		},
		{
			"-.3",
			true,
		},
		{
			"2e10",
			true,
		},
		{
			"90E3",
			true,
		},
		{
			"3e+7",
			true,
		},
		{
			"+6e-1",
			true,
		},
		{
			"53.5e93",
			true,
		},
		{
			"-123.456e789",
			true,
		},
		{
			"abc",
			false,
		},
		{
			"1a",
			false,
		},
		{
			"1e",
			false,
		},
		{
			"e3",
			false,
		},
		{
			"99e2.5",
			false,
		},
		{
			"--6",
			false,
		},
		{
			"-+3",
			false,
		},
		{
			"95a54e53",
			false,
		},
		{
			"1.3e+3+",
			false,
		},
	}

	for _, t := range tests {
		if isNumber(t.s) != t.r {
			fmt.Printf("failed: %s, %s\n", t.s, t.r)
		} else {
			fmt.Printf("passed: %s, %v\n", t.s, t.r)
		}
	}
}