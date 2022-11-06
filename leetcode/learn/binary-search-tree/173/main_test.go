package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	root   *TreeNode
	expect []int
}{
	{
		root:   testCase(),
		expect: []int{1, 2, 3, 4, 5, 6},
	},
}

func testCase() *TreeNode {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 3}
	four := &TreeNode{Val: 4}
	five := &TreeNode{Val: 5}
	six := &TreeNode{Val: 6}

	four.Left = two
	four.Right = five

	two.Left = one
	two.Right = three

	five.Right = six

	return four
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			iterator := Constructor(testCase.root)
			var output []int
			for iterator.HasNext() {
				next := iterator.Next()
				output = append(output, next)
			}

			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Output is not equals to expected")
			}
		})
	}
}
