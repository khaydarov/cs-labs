package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	n    int
	want []int
}{
	{
		n:    0,
		want: []int{0},
	},
	{
		n:    1,
		want: []int{0, 1},
	},
	{
		n:    2,
		want: []int{0, 1, 3, 2},
	},
	{
		n:    3,
		want: []int{0, 1, 3, 2, 6, 7, 5, 4},
	},
}

func TestSolution1(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			result := grayCode1(tc.n)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("Expected %v but got %v", tc.want, result)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			result := grayCode2(tc.n)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("Expected %v but got %v", tc.want, result)
			}
		})
	}
}

func TestSolution3(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			result := grayCode3(tc.n)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("Expected %v but got %v", tc.want, result)
			}
		})
	}
}
