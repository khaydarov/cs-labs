package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	nums   []int
	k      int
	add    []int
	expect []int
}{
	{
		nums:   []int{4, 5, 8, 2},
		k:      3,
		add:    []int{3, 5, 10, 9, 4},
		expect: []int{4, 5, 5, 8, 8},
	},
	{
		nums:   []int{4, 5, 8},
		k:      1,
		add:    []int{6, 8, 9},
		expect: []int{8, 8, 9},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			structure := Constructor(testCase.k, testCase.nums)

			for key, val := range testCase.add {
				output := structure.Add(val)
				expect := testCase.expect[key]

				if output != expect {
					t.Errorf("output is not equals to expected")
				}
			}
		})
	}
}
