package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	zero := TreeNode{Val: 0}
	one := TreeNode{Val: 1}
	two := TreeNode{Val: 2}
	three := TreeNode{Val: 3}
	four := TreeNode{Val: 4}
	five := TreeNode{Val: 5}

	zero.Left = &one

	one.Left = &two
	one.Right = &three

	three.Left = &four
	three.Right = &five

	result := InorderTraversal(&zero)
	if fmt.Sprint(result) != "[2 1 4 3 5 0]" {
		t.Fatalf("failed test")
	}
}
