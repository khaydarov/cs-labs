package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	n      int
	k      int
	expect [][]int
}{
	{
		n: 4,
		k: 3,
		expect: [][]int{
			{1, 2, 3},
			{1, 2, 4},
			{1, 3, 4},
			{2, 3, 4},
		},
	},
	{
		n: 1,
		k: 1,
		expect: [][]int{
			{1},
		},
	},
	{
		n: 4,
		k: 1,
		expect: [][]int{
			{1},
			{2},
			{3},
			{4},
		},
	},
	{
		n: 4,
		k: 2,
		expect: [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			got := combine(tc.n, tc.k)
			if len(got) != len(tc.expect) {
				t.Fatalf("expect %v, got %v", tc.expect, got)
			}
			for j, v := range got {
				for k, vv := range v {
					if vv != tc.expect[j][k] {
						t.Fatalf("expect %v, got %v", tc.expect, got)
					}
				}
			}
		})
	}
}
