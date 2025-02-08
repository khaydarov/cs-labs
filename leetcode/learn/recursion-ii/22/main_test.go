package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	n    int
	want []string
}{
	{
		n:    3,
		want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
	},
	{
		n:    1,
		want: []string{"()"},
	},
	{
		n:    2,
		want: []string{"(())", "()()"},
	},
}

func TestSolution1(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			if got := generateParenthesis1(tc.n); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TestSolution(%d) = %v; want %v", i, got, tc.want)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			if got := generateParenthesis2(tc.n); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TestSolution(%d) = %v; want %v", i, got, tc.want)
			}
		})
	}
}

func TestSolution3(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			if got := generateParenthesis3(tc.n); !reflect.DeepEqual(got, tc.want) {
				// Any order is fine
				t.Errorf("TestSolution(%d) = %v; want %v", i, got, tc.want)
			}
		})
	}
}
