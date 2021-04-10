package main

import "fmt"

func main() {
	arr := []int{1, 2, 4, 4}
	a := sumN(arr, 2)
	fmt.Println(a)
}

func sumN(arr []int, n int) bool {
	m := make(map[int]bool, len(arr))

	for _, v := range arr {
		if m[n - v] {
			return true
		}
		m[v] = true
	}

	return false
}