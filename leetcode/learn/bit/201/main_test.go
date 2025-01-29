package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	left   int
	right  int
	result int
}{
	{
		left:   5,
		right:  7,
		result: 4,
	},
	{
		left:   0,
		right:  0,
		result: 0,
	},
	{
		left:   1,
		right:  2147483647,
		result: 0,
	},
}

func TestSolution1(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			ans := rangeBitwiseAnd1(testCase.left, testCase.right)
			if ans != testCase.result {
				t.Errorf("Expected %d but got %d", testCase.result, ans)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			ans := rangeBitwiseAnd2(testCase.left, testCase.right)
			if ans != testCase.result {
				t.Errorf("Expected %d but got %d", testCase.result, ans)
			}
		})
	}
}

func TestSolution3(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			ans := rangeBitwiseAnd3(testCase.left, testCase.right)
			if ans != testCase.result {
				t.Errorf("Expected %d but got %d", testCase.result, ans)
			}
		})
	}
}
