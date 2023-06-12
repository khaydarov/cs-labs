package main

const base = 1337

func superPow(a int, b []int) int {
	var helper func(cursor int) int
	helper = func(cursor int) int {
		if cursor == -1 {
			return 1
		}

		return powMod(a, b[cursor], base) * powMod(helper(cursor-1), 10, base) % base
	}

	return helper(len(b) - 1)
}

// TC: O(log N)
// SC: O(log N)
func powMod(a int, p int, base int) int {
	if p == 0 {
		return 1
	}

	if p == 1 {
		return a % base
	}

	if p%2 == 0 {
		result := powMod(a, p/2, base)

		return (result * result) % base
	}

	return (a * powMod(a, p-1, base)) % base
}
