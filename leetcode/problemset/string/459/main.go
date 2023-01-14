package main

func repeatedSubstringPattern(s string) bool {
	t := s + s
	t = t[1 : len(t)-1]

	for i := 0; i < len(t); i++ {
		if i+len(s) > len(t) {
			break
		}

		if t[i:i+len(s)] == s {
			return true
		}
	}
	return false
}
