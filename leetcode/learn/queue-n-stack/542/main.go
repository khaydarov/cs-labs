package main

type Position struct {
	x int
	y int
}

type Queue struct {
	items []Position
}

func (q *Queue) Push(p Position) {
	q.items = append(q.items, p)
}

func (q *Queue) Front() Position {
	if q.Empty() {
		return Position{}
	}

	return q.items[0]
}

func (q *Queue) Pop() Position {
	if q.Empty() {
		return Position{}
	}

	front := q.items[0]
	q.items = q.items[1:]

	return front
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

func updateMatrix(mat [][]int) [][]int {
	q := Queue{}

	maxValue := 100000
	var result [][]int
	for i := 0; i < len(mat); i++ {
		var row []int
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				row = append(row, 0)
				q.Push(Position{i, j})
			} else {
				row = append(row, maxValue)
			}
		}

		result = append(result, row)
	}

	dirs := [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
	}

	for !q.Empty() {
		current := q.Pop()

		for _, dir := range dirs {
			nextX := current.x + dir[0]
			nextY := current.y + dir[1]

			if nextX >= 0 && nextX < len(mat) && nextY >= 0 && nextY < len(mat[0]) {
				if result[nextX][nextY] > result[current.x][current.y]+1 {
					result[nextX][nextY] = result[current.x][current.y] + 1
					q.Push(Position{nextX, nextY})
				}
			}
		}
	}

	return result
}
