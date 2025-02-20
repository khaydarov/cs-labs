package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	matrix [][]int
	want   [][]int
}{
	{
		matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		want: [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		},
	},
	{
		matrix: [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		},
		want: [][]int{
			{15, 13, 2, 5},
			{14, 3, 4, 1},
			{12, 6, 8, 9},
			{16, 7, 10, 11},
		},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			rotate(tc.matrix)
			for i := range tc.matrix {
				for j := range tc.matrix[i] {
					if tc.matrix[i][j] != tc.want[i][j] {
						t.Fatalf("got: %v, want: %v", tc.matrix, tc.want)
					}
				}
			}
		})
	}
}
