package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	dictionary []string
	search     []string
	expect     []bool
}{
	{
		dictionary: []string{"bad", "dad", "mad"},
		search:     []string{"pad", "bad", ".ad", "b.."},
		expect:     []bool{false, true, true, true},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			w := Constructor()
			for _, word := range tc.dictionary {
				w.AddWord(word)
			}

			for j, s := range tc.search {
				if got, want := w.Search(s), tc.expect[j]; got != want {
					t.Errorf("got %v, want %v", got, want)
				}
			}
		})
	}
}
