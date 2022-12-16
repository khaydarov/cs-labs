package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	dictionary []string
	words      []string
	expect     []bool
}{
	{
		[]string{"door", "dear", "cake", "cook"},
		[]string{"dear", "coke", "cake"},
		[]bool{false, false, true},
	},
	{
		[]string{"deep", "duke", "cube", "cool", "happy"},
		[]string{"dear", "deep", "disk", "cute", "honor"},
		[]bool{true, true, true, false, true},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			c := Constructor(testCase.dictionary)
			var output []bool
			for _, word := range testCase.words {
				output = append(output, c.IsUnique(word))
			}

			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Wrong answer. Expected: %v, got %v", testCase.expect, output)
			}
		})
	}

}
