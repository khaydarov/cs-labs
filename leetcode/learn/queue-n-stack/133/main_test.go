package main

import (
	"reflect"
	"testing"
)

func testCase() *Node {
	one := Node{1, nil}
	two := Node{2, nil}
	three := Node{3, nil}
	four := Node{4, nil}

	one.Neighbors = []*Node{&two, &four}
	two.Neighbors = []*Node{&one, &three}
	three.Neighbors = []*Node{&two, &four}
	four.Neighbors = []*Node{&one, &three}

	return &one
}

func TestSolution(t *testing.T) {
	output := cloneGraph(testCase())
	if !reflect.DeepEqual(output, testCase()) {
		t.Errorf("Wrong answer")
	}
}
