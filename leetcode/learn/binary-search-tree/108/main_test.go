package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	nums   []int
	expect *TreeNode
}{
	{
		[]int{1, 3},
		testCase(),
	},
}

func testCase() *TreeNode {
	one := &TreeNode{Val: 3}
	two := &TreeNode{Val: 1}

	one.Left = two

	return one
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := sortedArrayToBST(testCase.nums)
			if !reflect.DeepEqual(output, testCase.expect) {
				t.Errorf("Output does not equals to expected")
			}
		})
	}
}
