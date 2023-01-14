package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	strings []string
	expect  [][]string
}{
	{
		[]string{"remain", "marine"},
		[][]string{
			{"remain", "marine"},
		},
	},
	{
		[]string{"abc", "bca", "sort", "tors"},
		[][]string{
			{"abc", "bca"},
			{"sort", "tors"},
		},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := groupAnagrams(testCase.strings)
			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Wrong answer. Expected %v, but got %v", testCase.expect, output)
			}
		})
	}
}
