package main

func isValid(s string) bool {
	stack := NewStack()

	for _, c := range s {
		if c == '(' {
			stack.Push(c)
		} else {
			if stack.IsEmpty() {
				return false
			}

			stack.Pop()
		}
	}

	return stack.IsEmpty()
}

// TC: O(2^2n * n)
// SC: O(2^2n * n)
func generateParenthesis1(n int) []string {
	var result []string

	q := NewQueue()
	q.Enqueue("(")

	for !q.IsEmpty() {
		element := q.Dequeue()
		if len(element) == n*2 {
			if isValid(element) {
				result = append(result, element)
			}
			continue
		}

		q.Enqueue(element + "(")
		q.Enqueue(element + ")")
	}

	return result
}

// TC: O(4^N / sqrt(N))
// SC: O(N)
func generateParenthesis2(n int) []string {
	var result []string

	var backtrack func(s string, forward, backward int)
	backtrack = func(s string, forward, backward int) {
		if forward == 0 && backward == 0 {
			result = append(result, s)
			return
		}

		if forward > 0 {
			backtrack(s+"(", forward-1, backward)
		}

		if backward > 0 && forward < backward {
			backtrack(s+")", forward, backward-1)
		}
	}

	backtrack("", n, n)

	return result
}

func generateParenthesis3(n int) []string {
	var memo = make(map[int][]string)
	var helper func(n int, cache map[int][]string) []string
	helper = func(n int, cache map[int][]string) []string {
		if n == 0 {
			return []string{""}
		}

		if v, ok := cache[n]; ok {
			return v
		}

		var result []string
		for i := 0; i < n; i++ {
			for _, left := range helper(i, cache) {
				for _, right := range helper(n-i-1, cache) {
					result = append(result, "("+left+")"+right)
				}
			}
		}

		cache[n] = result
		return result
	}

	return helper(n, memo)
}
