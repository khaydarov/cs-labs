package main

import "math"

type ArrayReader struct {
}

func (this *ArrayReader) get(index int) int { return 0 }

func search(reader ArrayReader, target int) int {
	l := 0
	r := 9999

	max := math.MaxInt32

	for l <= r {
		pivot := l + (r-l)/2
		value := reader.get(pivot)

		// fmt.Println(pivot, value, l, r)
		if value == target {
			return pivot
		}

		if value == max || value > target {
			r = pivot - 1
		} else {
			l = pivot + 1
		}
	}

	return -1
}
