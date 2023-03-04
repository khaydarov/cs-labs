package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	inorder           []int
	postorder         []int
	expectedRootValue int
}{
	{
		[]int{9, 3, 15, 20, 7},
		[]int{9, 15, 7, 20, 3},
		3,
	},
	{
		[]int{-1},
		[]int{-1},
		-1,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			outputNode := buildTree(testCase.inorder, testCase.postorder)
			if outputNode.Val != testCase.expectedRootValue {
				t.Errorf("Wrong answer")
			}
		})
	}
}
