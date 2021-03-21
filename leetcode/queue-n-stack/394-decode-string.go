package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	data []byte
}

func (s *Stack) Push(x byte) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() byte {
	l := len(s.data)
	v := s.data[l - 1]

	s.data = s.data[:l - 1]

	return v
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Top() byte {
	l := len(s.data)

	return s.data[l - 1]
}

type NumStack struct {
	data []int
}

func (s *NumStack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *NumStack) Pop() int {
	l := len(s.data)
	v := s.data[l - 1]

	s.data = s.data[:l - 1]

	return v
}

func (s *NumStack) Top() int {
	l := len(s.data)

	return s.data[l - 1]
}

func reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}

	return result
}

func decodeString(s string) string {
	symbols := Stack{}
	numbers := NumStack{}

	number := ""
	for _, l := range s {
		if l >= '0' && l <= '9' {
			number += string(l)
		} else if l == '[' {
			m, _ := strconv.Atoi(number)
			numbers.Push(m)
			symbols.Push(byte(l))
			number = ""
		} else if l == ']' {
			str := ""
			for symbols.Top() != '[' {
				str += string(symbols.Pop())
			}
			symbols.Pop()

			str = reverse(str)
			k := numbers.Pop()

			repeated := ""
			for i := 1; i <= k; i++ {
				repeated += str
			}

			for _, strL := range repeated {
				symbols.Push(byte(strL))
			}
		} else {
			symbols.Push(byte(l))
		}
	}

	result := ""
	for !symbols.Empty() {
		result += string(symbols.Pop())
	}

	return reverse(result)
}

func main() {
	//s := "3[a]2[bc]"
	//s := "3[a2[c]]"
	s := "2[abc]3[cd]ef"
	//s := "abc3[cd]xyz"
	//s := "11[x]c"
	fmt.Println(decodeString(s))
}