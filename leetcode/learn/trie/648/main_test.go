package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	dictionary []string
	sentence   string
	expected   string
}{
	{
		dictionary: []string{"cat", "bat", "rat"},
		sentence:   "the cattle was rattled by the battery",
		expected:   "the cat was rat by the bat",
	},
	{
		dictionary: []string{"a", "aa", "aaa", "aaaa"},
		sentence:   "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa",
		expected:   "a a a a a a a a bbb baba a",
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			actual := replaceWords(tc.dictionary, tc.sentence)
			if actual != tc.expected {
				t.Errorf("expected: %s, got: %s", tc.expected, actual)
			}
		})
	}
}
