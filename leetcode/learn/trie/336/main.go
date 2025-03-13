package main

import "strings"

func allValidPrefixes(word string) []string {
	var result []string
	for i := range word {
		if isPalindrome(word, i, len(word)-1) {
			result = append(result, word[0:i])
		}
	}
	return result
}

func allValidSuffixes(word string) []string {
	var result []string
	for i := range word {
		if isPalindrome(word, 0, i) {
			result = append(result, word[i+1:])
		}
	}
	return result
}

func palindromePairs(words []string) [][]int {
	var result [][]int

	wordsMap := make(map[string]int)
	for i, word := range words {
		wordsMap[word] = i
	}

	for i, word := range words {
		reverseWord := reverse(word)
		if v, ok := wordsMap[reverseWord]; ok && v != i {
			result = append(result, []int{i, v})
		}

		for _, prefix := range allValidPrefixes(word) {
			reversePrefix := reverse(prefix)
			if v, ok := wordsMap[reversePrefix]; ok && v != i {
				result = append(result, []int{i, v})
			}
		}

		for _, suffix := range allValidSuffixes(word) {
			reverseSuffix := reverse(suffix)
			if v, ok := wordsMap[reverseSuffix]; ok && v != i {
				result = append(result, []int{v, i})
			}
		}
	}

	return result
}

func reverse(word string) string {
	var reversed strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		reversed.WriteByte(word[i])
	}
	return reversed.String()
}

func isPalindrome(target string, from, to int) bool {
	for from <= to {
		if target[from] != target[to] {
			return false
		}

		from++
		to--
	}

	return true
}
