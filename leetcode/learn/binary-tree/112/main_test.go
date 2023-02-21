package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	root      *TreeNode
	targetSum int
	expected  bool
}{
	{
		testCase1(),
		6,
		true,
	},
	{
		testCase2(),
		9,
		true,
	},
}

func testCase1() *TreeNode {
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{2, nil, nil}
	three := &TreeNode{2, nil, nil}
	four := &TreeNode{3, nil, nil}
	five := &TreeNode{3, nil, nil}
	six := &TreeNode{4, nil, nil}

	one.Left = two
	one.Right = three

	two.Right = four
	three.Left = five

	five.Left = six

	return one
}

func testCase2() *TreeNode {

	one := &TreeNode{1, nil, nil}
	two := &TreeNode{2, nil, nil}
	three := &TreeNode{2, nil, nil}
	four := &TreeNode{3, nil, nil}
	five := &TreeNode{4, nil, nil}
	six := &TreeNode{6, nil, nil}
	seven := &TreeNode{1, nil, nil}

	one.Left = two
	one.Right = three

	two.Left = four
	two.Right = five

	three.Left = six
	three.Right = seven

	return one
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := hasPathSum(testCase.root, testCase.targetSum)

			if output != testCase.expected {
				t.Errorf("Wrong answer. Expecter %v but got %v", testCase.expected, output)
			}
		})
	}
}
