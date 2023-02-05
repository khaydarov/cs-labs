package main

/*
*
-- [1, 2, 3, 4, 5, 6, 7]

[5, 6, 7, 1, 2, 3, 4]

Two sorted arrays: F and S;

1) if mid is in F and target is in F. Then we search in [start, mid]; Compare mid and target;
2) if mid is in F and target is in S. Then we search in [mid, end];
3) if mid is in S and target is in F. Then we search in [start, mid];
4) if mid is in S and target is in S. Then we search in [mid, end]; Compare mid and target;
*/
func search(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return true
		}

		if !isBinarySearchHelpful(nums, l, nums[mid]) {
			l++
			continue
		}

		midInFirst := existInFirst(nums, l, nums[mid])
		targetInFirst := existInFirst(nums, l, target)

		// fmt.Println(midInFirst, targetInFirst, nums[mid], mid, l, r)

		if midInFirst && targetInFirst {
			if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}

		} else if midInFirst && !targetInFirst {
			l = mid + 1
		} else if !midInFirst && targetInFirst {
			r = mid - 1
		} else {
			if target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}

func existInFirst(nums []int, start, value int) bool {
	return nums[start] <= value
}

func xor(a, b bool) bool {
	return (a || b) && !(a && b)
}

func isBinarySearchHelpful(nums []int, start, value int) bool {
	return nums[start] != value
}
