package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	n      int
	result []int
}{
	{
		n:      1,
		result: []int{0, 1},
	},
	{
		n:      2,
		result: []int{0, 1, 1},
	},
	{
		n:      3,
		result: []int{0, 1, 1, 2},
	},
	{
		n:      5,
		result: []int{0, 1, 1, 2, 1, 2},
	},
}

func TestSolution1(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			result := countBits1(testCase.n)
			if !reflect.DeepEqual(result, testCase.result) {
				t.Errorf("Extected %v but got %v", testCase.result, result)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			result := countBits2(testCase.n)
			if !reflect.DeepEqual(result, testCase.result) {
				t.Errorf("Extected %v but got %v", testCase.result, result)
			}
		})
	}
}

func TestSolution3(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			result := countBits3(testCase.n)
			if !reflect.DeepEqual(result, testCase.result) {
				t.Errorf("Extected %v but got %v", testCase.result, result)
			}
		})
	}
}
