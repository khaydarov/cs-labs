package main

import (
	"fmt"
	"testing"
)

func TestRecursiveSolution(t *testing.T) {
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

	result := PreorderTraversalRecursive(&zero)
	if fmt.Sprint(result) != "[0 1 2 3 4 5]" {
		t.Fatalf("failed test")
	}
}

func TestIterativeSolution(t *testing.T) {
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

	result := PreorderTraversalIterative(&zero)
	if fmt.Sprint(result) != "[0 1 2 3 4 5]" {
		t.Fatalf("failed test")
	}
}
