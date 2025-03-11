package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	board [][]byte
	words []string
	want  []string
}{
	{
		board: [][]byte{
			{'o', 'a', 'a', 'n'},
			{'e', 't', 'a', 'e'},
			{'i', 'h', 'k', 'r'},
			{'i', 'f', 'l', 'v'},
		},
		words: []string{"oath", "pea", "eat", "rain"},
		want:  []string{"oath", "eat"},
	},
	{
		board: [][]byte{
			{'a', 'b'},
			{'c', 'd'},
		},
		words: []string{"abdca", "acd"},
		want:  []string{"acd"},
	},
	{
		board: [][]byte{
			{'a', 'b', 'c'},
			{'a', 'e', 'd'},
			{'a', 'f', 'g'},
		},
		words: []string{"eaafgdcba"},
		want:  []string{"eaafgdcba"},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			result := findWords(tc.board, tc.words)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("Wrong answer. Expected %v, got %v", tc.want, result)
			}
		})
	}
}
