package main

import "fmt"

func main() {

	arr := []int{1, 4, 2, 5, 3}

	fmt.Println(sumOddLengthSubarrays(arr))
}

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)

	if n == 1 {
		return arr[0]
	}

	for i := 1; i < n; i++ {
		arr[i] += arr[i - 1]
	}

	sum := 0
	for i := 3; i < n; i += 2 {
		sum += arr[i - 1]
		for j := 1; j <= n - i; j++ {
			sum += arr[j + i - 1] - arr[j - 1]
		}
	}

	if n % 2 == 0 {
		sum += arr[n - 1]
	} else {
		sum += 2 * arr[n - 1]
	}

	//fmt.Println(sum)
	return sum
}
