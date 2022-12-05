package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums   []int
	expect int
}{
	{
		[]int{1, 1, 1, 0, 0, 1, 1},
		4,
	},
	{
		[]int{1, 0, 0, 1, 1, 0, 1},
		4,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			output := findMaxConsecutiveOnes(testCase.nums)

			if output != testCase.expect {
				t.Errorf("Wrong answer")
			}
		})
	}
}

func TestStreamSolution(t *testing.T) {
	s := Stream{}
	fmt.Println(s.Add(1))
	fmt.Println(s.Add(1))
	fmt.Println(s.Add(0))
	fmt.Println(s.Add(0))
	fmt.Println(s.Add(1))
	fmt.Println(s.Add(1))
	fmt.Println(s.Add(1))
	fmt.Println(s.Add(1))

}
