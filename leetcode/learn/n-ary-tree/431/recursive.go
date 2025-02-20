package main

func recursiveEncoder(root *Node) *TreeNode {
	if root == nil {
		return nil
	}

	treeNode := &TreeNode{Val: root.Val}
	if len(root.Children) > 0 {
		firstNode := root.Children[0]
		treeNode.Left = recursiveEncoder(firstNode)
	}

	sibling := treeNode.Left
	for i := 1; i < len(root.Children); i++ {
		sibling.Right = recursiveEncoder(root.Children[i])
		sibling = sibling.Right
	}

	return treeNode
}

func recursiveDecoder(root *TreeNode) *Node {
	if root == nil {
		return nil
	}

	node := &Node{Val: root.Val}
	sibling := root.Left
	for sibling != nil {
		node.Children = append(node.Children, recursiveDecoder(sibling))
		sibling = sibling.Right
	}

	return node
}
