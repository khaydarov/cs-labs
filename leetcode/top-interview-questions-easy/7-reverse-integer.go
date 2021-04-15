package main

import "fmt"

// TC: O(X), where X is the number of digits
// SC: O(1)
func reverse(x int) int {
	min := -2147483648
	max := 2147483647

	sum := 0
	for x != 0 {
		k := x % 10
		x /= 10

		if sum > max / 10 || (sum == max / 10 && k > 7)  {
			return 0
		}

		if sum < min / 10 || (sum == min / 10 && k < -8) {
			return 0
		}

		sum = sum * 10 + k
	}

	return sum
}

func main() {
	fmt.Println(reverse(-74))
}