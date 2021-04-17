package main

import "fmt"

func romanToInt(s string) int {
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]

		if i + 1 < len(s) && m[char] < m[s[i + 1]] {
			result -= m[char]
		} else {
			result += m[char]
		}
	}

	return result
}

func main() {
	tests := []struct{
		s string
		e int
	}{
		{"M", 1000},
		{"CM", 900},
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MCDDD", 2400},
		{"XIV", 14},
	}

	for _, test := range tests {
		r := romanToInt(test.s)
		if r == test.e {
			fmt.Println("pass", test.e)
		} else {
			fmt.Println("fail", test.e, r)
		}
	}
}