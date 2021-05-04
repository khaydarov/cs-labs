package main

type TreeNode struct {
	Val 		int
	Left, Right *TreeNode
}

type Queue struct {
	data []*TreeNode
}

func (q *Queue) Enqueue(node *TreeNode) {
	q.data = append(q.data, node)
}

func (q *Queue) Dequeue() *TreeNode {
	if len(q.data) == 0 {
		return nil
	}

	node := q.data[0]
	q.data = q.data[1:]

	return node
}

func (q *Queue) Length() int {
	return len(q.data)
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	q := Queue{}
	q.Enqueue(root)

	for !q.Empty() {
		l := q.Length()

		var level []int
		for i := 0; i < l; i++ {
			node := q.Dequeue()

			if node.Left != nil {
				q.Enqueue(node.Left)
			}

			if node.Right != nil {
				q.Enqueue(node.Right)
			}
			level = append(level, node.Val)
		}
		result = append(result, level)
	}

	return result
}

func main() {
	//one := &TreeNode{1, nil, nil}
	//two := &TreeNode{2, nil, nil}
	//three := &TreeNode{3, nil, nil}
	//four := &TreeNode{4, nil, nil}
	//five := &TreeNode{5, nil, nil}
	//six := &TreeNode{6, nil, nil}
	//seven := &TreeNode{7, nil, nil}
	//eight := &TreeNode{8, nil, nil}
	//
	//one.Left = two
	//one.Right = three
	//
	//two.Left = four
	//two.Right = five
	//
	//five.Left = six
	//five.Right = seven
	//
	//three.Right = eight

	//one := &TreeNode{1, nil, nil}
	//two := &TreeNode{4, nil, nil}
	//three := &TreeNode{2, nil, nil}
	//four := &TreeNode{6, nil, nil}
	//five := &TreeNode{7, nil, nil}
	//six := &TreeNode{5, nil, nil}
	//seven := &TreeNode{3, nil, nil}
	//eight := &TreeNode{9, nil, nil}
	//
	//one.Left = two
	//one.Right =  three
	//
	//two.Left = four
	//two.Right = five
	//
	//three.Left = six
	//three.Right = seven
	//
	//four.Left = eight
}