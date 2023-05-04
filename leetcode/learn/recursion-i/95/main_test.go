package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	n      int
	expect []string
}{
	{
		1,
		[]string{"1##"},
	},
	{
		2,
		[]string{"1#2##", "21###"},
	},
	{
		3,
		[]string{"1#2#3##", "1#32###", "21##3##", "321####", "31#2###"},
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			trees := generateTrees(testCase.n)

			var serializedTrees []string
			for _, tree := range trees {
				serializedTrees = append(serializedTrees, serialize(tree))
			}

			if !reflect.DeepEqual(serializedTrees, testCase.expect) {
				t.Errorf("Wrong answer. Expected %v got %v", testCase.expect, serializedTrees)
			}
		})
	}
}
