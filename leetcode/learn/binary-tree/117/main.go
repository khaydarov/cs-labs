package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Queue struct {
	data []*Node
}

func (q *Queue) Length() int {
	return len(q.data)
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue) Front() *Node {
	if len(q.data) == 0 {
		return nil
	}

	return q.data[0]
}

func (q *Queue) Enqueue(node *Node) {
	q.data = append(q.data, node)
}

func (q *Queue) Dequeue() *Node {
	if len(q.data) == 0 {
		return nil
	}

	front := q.Front()
	q.data = q.data[1:]

	return front
}

// TC: O(N)
// SC: O(N)
func connect1(root *Node) *Node {
	if root == nil {
		return root
	}

	q := Queue{}
	q.Enqueue(root)

	for !q.Empty() {
		l := q.Length()

		nodesToConnect := Queue{}
		for i := 0; i < l; i++ {
			current := q.Dequeue()

			if current.Left != nil {
				q.Enqueue(current.Left)
			}

			if current.Right != nil {
				q.Enqueue(current.Right)
			}

			nodesToConnect.Enqueue(current)
		}

		for !nodesToConnect.Empty() {
			current := nodesToConnect.Dequeue()
			current.Next = nodesToConnect.Front()
		}
	}

	return root
}

// TC: O(N)
// SC: O(1)
func connect(root *Node) *Node {
	var leftMost *Node
	leftMost = root

	for leftMost != nil {
		var head *Node
		var prev *Node

		head = leftMost
		leftMost = nil

		for head != nil {
			if head.Left != nil {
				if leftMost == nil {
					leftMost = head.Left
				}

				if prev == nil {
					prev = head.Left
				} else {
					prev.Next = head.Left
					prev = prev.Next
				}
			}

			if head.Right != nil {
				if leftMost == nil {
					leftMost = head.Right
				}

				if prev == nil {
					prev = head.Right
				} else {
					prev.Next = head.Right
					prev = prev.Next
				}
			}

			head = head.Next
		}
	}

	return root
}
