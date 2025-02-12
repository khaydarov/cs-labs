package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	graph := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{3, 5},
		{3, 6},
		{4, 7},
		{6, 7},
	}

	r := BFS(graph, 1, 6)
	if r != 3 {
		t.Error("Expected 3, got ", r)
	}
}
