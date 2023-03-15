package main

import (
	"reflect"
	"testing"
)

func testCase1() (*Node, *Node) {
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}

	two.Left = one
	one.Parent = two

	two.Right = three
	three.Parent = two

	return one, two
}

func testCase2() (*Node, *Node) {
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}
	four := &Node{Val: 4}
	five := &Node{Val: 5}
	six := &Node{Val: 6}

	five.Left = three
	three.Parent = five

	five.Right = six
	six.Parent = five

	three.Left = two
	two.Parent = three

	three.Right = four
	four.Parent = three

	two.Left = one
	one.Parent = two

	return four, five
}

func testCase3() (*Node, *Node) {
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}
	four := &Node{Val: 4}
	five := &Node{Val: 5}
	six := &Node{Val: 6}

	five.Left = three
	three.Parent = five

	five.Right = six
	six.Parent = five

	three.Left = two
	two.Parent = three

	three.Right = four
	four.Parent = three

	two.Left = one
	one.Parent = two

	return six, nil
}

func TestSolution1(t *testing.T) {
	var node, successor, output *Node
	node, successor = testCase1()

	output = inorderSuccessor(node)
	if !reflect.DeepEqual(output, successor) {
		t.Errorf("Wrong answer. Expected %v but got %v", successor, output)
	}

	node, successor = testCase2()
	output = inorderSuccessor(node)
	if !reflect.DeepEqual(output, successor) {
		t.Errorf("Wrong answer. Expect %v but got %v", successor, output)
	}

	node, successor = testCase3()
	output = inorderSuccessor(node)
	if !reflect.DeepEqual(output, successor) {
		t.Errorf("Wrong answer. Expect %v but got %v", successor, output)
	}
}
