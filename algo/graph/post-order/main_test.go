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

	result := PostOrderTravelsal(&zero)
	if fmt.Sprint(result) != "[2 4 5 3 1 0]" {
		t.Fatalf("failed test")
	}
}
