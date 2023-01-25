package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	temperatures []int
	expected     []int
}{
	{
		[]int{73, 74, 75, 71, 69, 72, 76, 73},
		[]int{1, 1, 4, 2, 1, 1, 0, 0},
	},
	{
		[]int{30, 40, 50, 60},
		[]int{1, 1, 1, 0},
	},
	{
		[]int{73},
		[]int{0},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := dailyTemperatures(testCase.temperatures)
			if reflect.DeepEqual(output, testCase.expected) {
				t.Errorf("Wrong answer. Extected %v, got %v", testCase.expected, output)
			}
		})
	}
}
