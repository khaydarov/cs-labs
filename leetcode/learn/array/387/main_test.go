package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	givenString string
	expect      int
}{
	{
		"loveleetcode",
		2,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := firstUniqChar(testCase.givenString)
			if output != testCase.expect {
				t.Errorf("Not equal to expected %d", testCase.expect)
			}
		})
	}
}
