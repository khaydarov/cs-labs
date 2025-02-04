package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	graph [][]int
	x     int
	y     int
	want  bool
}{
	{
		graph: [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 5},
			{3, 5},
			{3, 6},
			{4, 7},
			{6, 7},
		},
		x:    5,
		y:    5,
		want: true,
	},
	{
		graph: [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 5},
			{3, 5},
			{3, 6},
			{4, 7},
			{6, 7},
		},
		x:    1,
		y:    7,
		want: true,
	},
	{
		graph: [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 5},
			{3, 5},
			{3, 6},
			{4, 7},
			{6, 7},
		},
		x:    4,
		y:    5,
		want: false,
	},
}

func TestRecursive(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestRecursive-%d", i), func(t *testing.T) {
			got := Recursive(tc.graph, tc.x, tc.y)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func TestIterative(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("TestIterative-%d", i), func(t *testing.T) {
			got := Iterative(tc.graph, tc.x, tc.y)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
