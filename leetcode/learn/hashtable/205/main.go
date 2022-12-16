package main

func isIsomorphic(s string, t string) bool {
	transformationsST := make(map[byte]byte)
	transformationsTS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		_, sExists := transformationsST[s[i]]
		_, tExists := transformationsTS[t[i]]

		if (!sExists && tExists && s[i] != transformationsTS[t[i]]) || (sExists && transformationsST[s[i]] != t[i]) {
			return false
		}

		transformationsST[s[i]] = t[i]
		transformationsTS[t[i]] = s[i]
	}

	result := ""
	for i := 0; i < len(s); i++ {
		result += string(transformationsST[s[i]])
	}

	if result != t {
		return false
	}

	return true
}
