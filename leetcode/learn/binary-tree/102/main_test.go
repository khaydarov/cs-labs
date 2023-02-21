package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	root   *TreeNode
	expect [][]int
}{
	{
		testCase1(),
		[][]int{
			{1},
			{2, 3},
			{4, 5, 8},
			{6, 7},
		},
	},
	{
		testCase2(),
		[][]int{
			{1},
			{4, 2},
			{6, 7, 5, 3},
			{9},
		},
	},
}

func testCase1() *TreeNode {
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{2, nil, nil}
	three := &TreeNode{3, nil, nil}
	four := &TreeNode{4, nil, nil}
	five := &TreeNode{5, nil, nil}
	six := &TreeNode{6, nil, nil}
	seven := &TreeNode{7, nil, nil}
	eight := &TreeNode{8, nil, nil}

	one.Left = two
	one.Right = three

	two.Left = four
	two.Right = five

	five.Left = six
	five.Right = seven

	three.Right = eight

	return one
}

func testCase2() *TreeNode {
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

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := levelOrder(testCase.root)

			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, output)
			}
		})
	}
}
