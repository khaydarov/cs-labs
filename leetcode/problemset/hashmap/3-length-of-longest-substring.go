package hashmap

// TC: O(N^2)
// SC: O(N)
func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		m := make(map[byte]bool)
		t := ""
		l := 0
		for j := i; j < len(s) && m[s[j]] != true; j++ {
			l++
			m[s[j]] = true
			t += string(s[j])
		}

		if max < l {
			max = l
		}
	}

	return max
}

// TC: O(N)
// SC: O(N)
func lengthOfLongestSubstring1(s string) int {
	start := 0
	end := 0
	max := 0
	m := make(map[byte]int)
	for end < len(s) {
		if _, ok := m[s[end]]; ok {
			if m[s[end]] + 1 > start {
				start = m[s[end]] + 1
			}
		}

		m[s[end]] = end
		if max < end - start + 1 {
			max = end - start + 1
		}
		end++
	}

	return max
}