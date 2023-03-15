package main

import (
	"reflect"
	"testing"
)

func testCase1() (*TreeNode, *TreeNode, *TreeNode) {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 3}

	two.Left = one
	two.Right = three

	return two, one, two
}

func testCase2() (*TreeNode, *TreeNode, *TreeNode) {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 3}
	four := &TreeNode{Val: 4}
	five := &TreeNode{Val: 5}
	six := &TreeNode{Val: 6}

	five.Left = three
	five.Right = six

	three.Left = two
	three.Right = four

	two.Left = one

	return five, four, five
}

func testCase3() (*TreeNode, *TreeNode, *TreeNode) {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	three := &TreeNode{Val: 3}
	four := &TreeNode{Val: 4}
	five := &TreeNode{Val: 5}
	six := &TreeNode{Val: 6}

	five.Left = three
	five.Right = six

	three.Left = two
	three.Right = four

	two.Left = one

	return five, six, nil
}

func TestSolution1(t *testing.T) {
	var root, p, successor, output *TreeNode
	root, p, successor = testCase1()

	output = inorderSuccessor1(root, p)
	if !reflect.DeepEqual(output, successor) {
		t.Errorf("Wrong answer. Expected %v but got %v", successor, output)
	}

	root, p, successor = testCase2()
	output = inorderSuccessor1(root, p)
	if !reflect.DeepEqual(output, successor) {
		t.Errorf("Wrong answer. Expect %v but got %v", successor, output)
	}

	root, p, successor = testCase3()
	output = inorderSuccessor1(root, p)
	if !reflect.DeepEqual(output, successor) {
		t.Errorf("Wrong answer. Expect %v but got %v", successor, output)
	}
}

func TestSolution2(t *testing.T) {
	var root, p, successor, output *TreeNode
	root, p, successor = testCase1()

	output = inorderSuccessor2(root, p)
	if !reflect.DeepEqual(output, successor) {
		t.Errorf("Wrong answer. Expected %v but got %v", successor, output)
	}

	root, p, successor = testCase2()
	output = inorderSuccessor2(root, p)
	if !reflect.DeepEqual(output, successor) {
		t.Errorf("Wrong answer. Expect %v but got %v", successor, output)
	}

	root, p, successor = testCase3()
	output = inorderSuccessor2(root, p)
	if !reflect.DeepEqual(output, successor) {
		t.Errorf("Wrong answer. Expect %v but got %v", successor, output)
	}
}
