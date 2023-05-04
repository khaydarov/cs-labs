package main

// TC: O(NlogN)
// SC: O(NlogN)
func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	if n == 2 && k == 1 {
		return 0
	}

	if n == 2 && k == 2 {
		return 1
	}

	l := myPowIterative(2, n-1)
	if k <= l/2 {
		return kthGrammar(n-1, k)
	}

	k -= l / 2
	p := l / 2
	if k <= p/2 {
		return kthGrammar(n, p/2+k)
	} else {
		return kthGrammar(n, k-p/2)
	}
}

func myPowIterative(x int, n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	var result int
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

func kthGrammar1(n int, k int) int {
	if n == 1 {
		return 0
	}

	if k%2 == 0 {
		if kthGrammar(n-1, k/2) == 0 {
			return 1
		} else {
			return 0
		}
	} else {
		if kthGrammar(n-1, (k+1)/2) == 0 {
			return 0
		} else {
			return 1
		}
	}
}
