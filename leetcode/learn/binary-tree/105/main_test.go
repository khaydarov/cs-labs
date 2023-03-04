package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	preorder          []int
	inorder           []int
	expectedRootValue int
}{
	{
		[]int{3, 9, 20, 15, 7},
		[]int{9, 3, 15, 20, 7},
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
			outputNode := buildTree(testCase.preorder, testCase.inorder)
			if outputNode.Val != testCase.expectedRootValue {
				t.Errorf("Wrong answer")
			}
		})
	}
}
