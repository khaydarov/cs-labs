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

	result := PreOrderRecursive(one)
	expected := []int{1, 2, 3, 5, 4, 6}
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

	result := PreOrderIterative(one)
	expected := []int{1, 2, 3, 5, 4, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
