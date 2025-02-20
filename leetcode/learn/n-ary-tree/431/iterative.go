package main

type Pair struct {
	nary   *Node
	binary *TreeNode
}

type Queue struct {
	data []*Pair
}

func (q *Queue) Enqueue(pair *Pair) {
	q.data = append(q.data, pair)
}

func (q *Queue) Dequeue() *Pair {
	if len(q.data) == 0 {
		return nil
	}

	pair := q.data[0]
	q.data = q.data[1:]

	return pair
}

func (q *Queue) Length() int {
	return len(q.data)
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func iterativeEncoder(root *Node) *TreeNode {
	if root == nil {
		return nil
	}

	binary := &TreeNode{}

	q := Queue{}
	q.Enqueue(&Pair{
		nary:   root,
		binary: binary,
	})

	for !q.Empty() {
		l := q.Length()

		for i := 0; i < l; i++ {
			pair := q.Dequeue()

			nary := pair.nary
			binary := pair.binary

			binary.Val = nary.Val
			if len(nary.Children) > 0 {
				binary.Left = &TreeNode{}
				firstNode := nary.Children[0]

				q.Enqueue(&Pair{
					nary:   firstNode,
					binary: binary.Left,
				})
			}

			sibling := binary.Left
			for j := 1; j < len(nary.Children); j++ {
				sibling.Right = &TreeNode{}
				q.Enqueue(&Pair{
					nary:   nary.Children[j],
					binary: sibling.Right,
				})

				sibling = sibling.Right
			}
		}
	}

	return binary
}

func iterativeDecoder(root *TreeNode) *Node {
	if root == nil {
		return nil
	}

	nary := &Node{Val: root.Val}

	q := Queue{}
	q.Enqueue(&Pair{
		nary:   nary,
		binary: root,
	})

	for !q.Empty() {
		l := q.Length()

		for i := 0; i < l; i++ {
			pair := q.Dequeue()

			nary := pair.nary
			binary := pair.binary

			sibling := binary.Left
			for sibling != nil {
				node := &Node{Val: sibling.Val}
				q.Enqueue(&Pair{
					nary:   node,
					binary: sibling,
				})

				nary.Children = append(nary.Children, node)
				sibling = sibling.Right
			}
		}
	}

	return nary
}
