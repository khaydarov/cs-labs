package main

import "fmt"

type TreeNode struct {
	Val 		int
	Left, Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return calcSum(root, root.Val, targetSum)
}

func calcSum(node *TreeNode, actualSum, targetSum int) bool {
	if isLeaf(node) {
		return actualSum == targetSum
	}

	left := false
	right := false

	if node.Left != nil {
		left = calcSum(node.Left, actualSum + node.Left.Val, targetSum)
	}

	if node.Right != nil {
		right = calcSum(node.Right, actualSum + node.Right.Val, targetSum)
	}

	return left || right
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}


func main() {
	//one := &TreeNode{1, nil, nil}
	//two := &TreeNode{2, nil, nil}
	//three := &TreeNode{2, nil, nil}
	//four := &TreeNode{3, nil, nil}
	//five := &TreeNode{3, nil, nil}
	////six := &TreeNode{4, nil, nil}
	//
	//one.Left = two
	//one.Right = three
	//
	//two.Right = four
	//three.Left = five

	//five.Left = six

	one := &TreeNode{1, nil, nil}
	two := &TreeNode{2, nil, nil}
	three := &TreeNode{2, nil, nil}
	four := &TreeNode{3, nil, nil}
	five := &TreeNode{4, nil, nil}
	six := &TreeNode{6, nil, nil}
	seven := &TreeNode{1, nil, nil}

	one.Left = two
	one.Right =  three

	two.Left = four
	two.Right = five

	three.Left = six
	three.Right = seven

	fmt.Println(hasPathSum(one, 9))
}