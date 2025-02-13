package main

import (
	"reflect"
	"testing"
)

func TestRecursiveSolution(t *testing.T) {
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}
	four := &Node{Val: 4}
	five := &Node{Val: 5}
	six := &Node{Val: 6}

	one.Children = []*Node{two, three, four}
	three.Children = []*Node{five}
	four.Children = []*Node{six}

	result := PostOrderRecursive(one)
	expected := []int{2, 5, 3, 6, 4, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestIterativeSolution(t *testing.T) {
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}
	four := &Node{Val: 4}
	five := &Node{Val: 5}
	six := &Node{Val: 6}

	one.Children = []*Node{two, three, four}
	three.Children = []*Node{five}
	four.Children = []*Node{six}

	result := PostOrderIterative(one)
	expected := []int{2, 5, 3, 6, 4, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
