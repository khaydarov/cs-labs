package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	tree   *TreeNode
	result []int
}{
	{
		tree:   testTree1(),
		result: []int{9, 6, 4, 7, 1, 5, 2, 3},
	},
	{
		tree:   testTree2(),
		result: []int{1, 3, 2},
	},
}

func testTree1() *TreeNode {
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{4, nil, nil}
	three := &TreeNode{2, nil, nil}
	four := &TreeNode{6, nil, nil}
	five := &TreeNode{7, nil, nil}
	six := &TreeNode{5, nil, nil}
	seven := &TreeNode{3, nil, nil}
	eight := &TreeNode{9, nil, nil}

	one.Left = two
	one.Right = three

	two.Left = four
	two.Right = five

	three.Left = six
	three.Right = seven

	four.Left = eight

	return one
}

func testTree2() *TreeNode {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 3}

	one.Right = two
	two.Left = three

	return one
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test:%d", i), func(t *testing.T) {
			output := inorderTraversal(testCase.tree)
			if !reflect.DeepEqual(output, testCase.result) {
				t.Errorf("Output is not equals to expected result: %v", testCase.result)
			}
		})
	}
}
