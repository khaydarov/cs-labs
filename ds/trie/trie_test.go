package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	commands []string
	values   []string
	expected []bool
}{
	{
		commands: []string{"Insert", "Search", "StartWith", "StartWith", "Insert", "Search"},
		values:   []string{"apple", "apple", "app", "app", "app", "app"},
		expected: []bool{true, true, true, true, true, true},
	},
	{
		commands: []string{"Insert", "Search", "StartWith", "Insert", "Search"},
		values:   []string{"apple", "app", "app", "app", "app"},
		expected: []bool{true, false, true, true, true},
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		trie := Construct()
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			for j, cmd := range tc.commands {
				switch cmd {
				case "Insert":
					trie.Insert(tc.values[j])
				case "Search":
					if got := trie.Search(tc.values[j]); got != tc.expected[j] {
						t.Errorf("Test case %v failed: expected %v but got %v", i, tc.expected[j], got)
					}
				case "StartWith":
					if got := trie.StartWith(tc.values[j]); got != tc.expected[j] {
						t.Errorf("Test case %v failed: expected %v but got %v", i, tc.expected[j], got)
					}
				}
			}
		})
	}
}
