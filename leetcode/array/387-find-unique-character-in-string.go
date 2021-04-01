package main

import "fmt"

func firstUniqChar(s string) int {
	m := [26]int{}
	for _, v := range s {
		m[v - 'a']++
	}

	for i, v := range s {
		if m[v - 'a'] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}