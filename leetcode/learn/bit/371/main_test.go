package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	a   int
	b   int
	sum int
}{
	{
		a:   1,
		b:   2,
		sum: 3,
	},
	{
		a:   10,
		b:   25,
		sum: 35,
	},
	{
		a:   -10,
		b:   3,
		sum: -7,
	},
	{
		a:   77,
		b:   7,
		sum: 84,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			res := getSum(testCase.a, testCase.b)
			if res != testCase.sum {
				t.Errorf("Expected %d but got %d", testCase.sum, res)
			}
		})
	}
}
