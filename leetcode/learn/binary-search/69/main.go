package main

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	l := 0
	r := x / 2
	for l <= r {
		pivot := l + (r-l)/2

		if pivot*pivot == x {
			return pivot
		} else if pivot*pivot > x {
			r = pivot - 1
		} else {
			l = pivot + 1
		}
	}

	return r
}

func mySqrt1(x int) int {
	if x < 2 {
		return x
	}

	x0 := float64(x)
	x1 := float64(x0+float64(x)/x0) / 2.0

	for abs(x1-x0) >= 1 {
		x0 = x1
		x1 = (x0 + float64(x)/x0) / 2.0
	}

	return int(x1)
}

func abs(x float64) float64 {
	if x < 0 {
		return -1 * x
	}

	return x
}
