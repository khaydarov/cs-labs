package main

func twoSum(numbers []int, target int) []int {
	index1 := 0
	index2 := len(numbers) - 1

	for index1 < index2 {
		value := numbers[index1] + numbers[index2]

		if value == target {
			return []int{index1 + 1, index2 + 1}
		}

		if value > target {
			index2--
		} else {
			index1++
		}
	}

	return []int{index1 + 1, index2 + 1}
}
