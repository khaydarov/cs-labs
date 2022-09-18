package main

func myPowIterative(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		x = 1 / x
		n *= -1
	}
	var result float64
	result = 1
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}

		x *= x
		n /= 2
	}

	return result
}

func myPowRecursive(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n%2 == 0 {
		return myPowRecursive(x, n/2) * myPowRecursive(x, n/2)
	}

	return x * myPowRecursive(x, n-1)
}
