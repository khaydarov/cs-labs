package main

// sortArray implements HeapSort
// TC: O(NlogN)
// SC: O(logN)
func sortArray(nums []int) []int {
	length := len(nums)
	for i := length; i >= 0; i-- {
		maxHeapify(nums, length, i)
	}

	for i := length - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		maxHeapify(nums, i, 0)
	}

	return nums
}

func maxHeapify(nums []int, size, index int) {
	l := left(index)
	r := right(index)

	largest := index
	for l < size && nums[l] > nums[largest] {
		largest = l
	}

	if r < size && nums[r] > nums[largest] {
		largest = r
	}

	if largest != index {
		nums[index], nums[largest] = nums[largest], nums[index]
		maxHeapify(nums, size, largest)
	}
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
