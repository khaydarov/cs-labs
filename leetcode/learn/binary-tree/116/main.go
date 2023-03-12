package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// Approach 1
// TC: O(Log N * Log N)
// SC: O(Log N)
func connect1(root *Node) *Node {
	if root == nil {
		return root
	}

	left := root.Left
	right := root.Right

	for left != nil {
		left.Next = right
		left = left.Right
		right = right.Left
	}

	connect1(root.Left)
	connect1(root.Right)

	return root
}

// Approach 2
// TC: O(N)
// SC: O(log N)
func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	_connect(root.Left, root.Right)

	return root
}

func _connect(left, right *Node) {
	if left == nil {
		return
	}

	left.Next = right

	_connect(left.Left, left.Right)
	_connect(left.Right, right.Left)
	_connect(right.Left, right.Right)
}

// Approach 3: Level order traversal

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

func connect3(root *Node) *Node {
	if root == nil {
		return root
	}

	q := Queue{}
	q.Enqueue(root)

	for !q.Empty() {
		l := q.Length()

		var previous, current *Node
		for i := 1; i <= l-1; i++ {
			current = q.Dequeue()
			if current.Left != nil {
				q.Enqueue(current.Left)
			}

			if current.Right != nil {
				q.Enqueue(current.Right)
			}

			if previous != nil {
				previous.Next = current
			}
			previous = current
		}
	}

	return root
}

// Approach 4
// TC: O(N)
// SC: O(1)
func connect4(root *Node) *Node {
	if root == nil {
		return root
	}

	leftMost := root

	for leftMost.Left != nil {
		head := leftMost
		for head != nil {
			head.Left.Next = head.Right

			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}

			head = head.Next
		}

		leftMost = leftMost.Left
	}

	return root
}
