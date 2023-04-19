package main

func findKthLargest1(nums []int, k int) int {
	positives := make([]int, 10001)
	negatives := make([]int, 10001)

	for _, v := range nums {
		if v >= 0 {
			positives[v]++
		} else {
			negatives[-1*v]++
		}
	}

	// lookup among positive values
	for i := 10000; i >= 0; i-- {
		if positives[i] == 0 {
			continue
		}

		if k > positives[i] {
			k -= positives[i]
		} else {
			return i
		}
	}

	// lookup among negative values
	for i := 1; i <= 10000; i++ {
		if negatives[i] == 0 {
			continue
		}

		if k > negatives[i] {
			k -= negatives[i]
		} else {
			return -1 * i
		}
	}

	return 0
}

func findKthLargest(nums []int, k int) int {
	return QuickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func QuickSelect(nums []int, left, right, k int) int {
	pivot := partition(nums, left, right)

	if pivot == k {
		return nums[pivot]
	} else if pivot < k {
		return QuickSelect(nums, pivot+1, right, k)
	} else {
		return QuickSelect(nums, left, pivot-1, k)
	}
}

func partition(nums []int, left, right int) int {
	index := left + (right-left)/2
	nums[index], nums[right] = nums[right], nums[index]

	for i := left; i < right; i++ {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	return left
}
