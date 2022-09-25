package main

import (
	"fmt"
	"strconv"
)

// TC: O(N * K * Log K)
// SC: O(N * K)
func groupAnagrams1(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		c := []byte(str)
		qSort(c, 0, len(c) - 1)
		hash := string(c)
		m[hash] = append(m[hash], str)
	}

	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func qSort(s []byte, left, right int) {
	if left >= right {
		return
	}

	pivot := s[(left + right) / 2]
	index := partition(s, left, right, pivot)

	qSort(s, left, index - 1)
	qSort(s, index, right)
}

func partition(s []byte, left, right int, pivot byte) int {
	for left <= right {
		for s[left] < pivot {
			left++
		}

		for s[right] > pivot {
			right--
		}

		if left <= right {
			s[left], s[right] = s[right], s[left]
			left++
			right--
		}
	}

	return left
}

// TC: O(N * K)
// SC: O(N * K)
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		str := strs[i]

		var occurence [26]int
		for j := 0; j < len(str); j++ {
			occurence[str[j] - 'a']++
		}

		key := ""
		for j := 0; j < 26; j++ {
			key += "#"
			key += strconv.Itoa(occurence[j])
		}

		m[key] = append(m[key], str)
	}

	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func main() {
	strs := []string{"remain", "marine"}
	fmt.Println(groupAnagrams(strs))
}