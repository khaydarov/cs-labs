package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	preorder          []int
	postorder         []int
	expectedRootValue int
}{
	{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{2, 4, 6, 5, 3, 1},
		1,
	},
	{
		[]int{1},
		[]int{1},
		1,
	},
	{
		[]int{1, 2, 3},
		[]int{2, 3, 1},
		1,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			outputNode := constructFromPrePost(testCase.preorder, testCase.postorder)
			if outputNode.Val != testCase.expectedRootValue {
				t.Errorf("Wrong answer")
			}
		})
	}
}
