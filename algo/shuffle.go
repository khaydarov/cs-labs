package main

import (
	"fmt"
	"math/rand"
	"time"
)

func FisherYatesShuffle(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		j := rand.Intn(n - i) + i

		nums[i], nums[j] = nums[j], nums[i]
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	FisherYatesShuffle(nums)
	fmt.Println(nums)
}