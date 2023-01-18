package main

// IsMatch
// TC: O(N) where N is max of text and pattern
// SC: O(1)
func IsMatch(text string, pattern string) bool {
	first := 0
	second := 0
	for second < len(pattern) {
		/**
		 * Next of pattern symbol if asterisk - current pattern symbol matches text symbol, or it is point. Then we go next until it is matching
		 */
		if second+1 < len(pattern) && pattern[second+1] == '*' {
			for first < len(text) && (text[first] == pattern[second] || pattern[second] == '.') {
				first++
			}
			second += 2
		} else if first < len(text) && text[first] == pattern[second] || pattern[second] == '.' {
			first++
			second++
		} else {
			return false
		}
	}

	if first != len(text) {
		return false
	}

	return true
}
