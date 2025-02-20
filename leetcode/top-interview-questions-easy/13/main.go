package main

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

		if i+1 < len(s) && m[char] < m[s[i+1]] {
			result -= m[char]
		} else {
			result += m[char]
		}
	}

	return result
}
