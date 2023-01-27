package main

type Coord struct {
	x, y int
}

// Implementation with Stack
// ------------------------------
type Stack struct {
	data []Coord
}

func (s *Stack) Push(c Coord) {
	s.data = append(s.data, c)
}

func (s *Stack) Pop() Coord {
	last := len(s.data) - 1
	item := s.data[last]

	s.data = s.data[:last]

	return item
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Using BFS and Stack without recursion and visited array
func numIslands1(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])

	s := Stack{}
	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	islands := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				s.Push(Coord{i, j})
				islands++
			}

			for s.IsEmpty() == false {
				curr := s.Pop()

				grid[curr.x][curr.y] = '0'

				for _, v := range dirs {
					newX := curr.x + v[0]
					newY := curr.y + v[1]

					if newX >= 0 && newX < n && newY >= 0 && newY < m && grid[newX][newY] == '1' {
						s.Push(Coord{newX, newY})
					}
				}
			}
		}
	}

	return islands
}

// Implementation with Queue
// -------------------------------------
type Queue struct {
	data []Coord
}

func (q *Queue) Enqueue(c Coord) {
	q.data = append(q.data, c)
}

func (q *Queue) Dequeue() {
	q.data = q.data[1:]
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue) Front() Coord {
	return q.data[0]
}

func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])

	var visited [300][300]byte

	bfs := func(i, j int) {
		q := Queue{}
		q.Enqueue(Coord{i, j})

		for !q.Empty() {
			curr := q.Front()
			visited[curr.x][curr.y] = 1
			dirs := [][]int{
				{-1, 0},
				{1, 0},
				{0, -1},
				{0, 1},
			}

			for _, v := range dirs {
				newX := curr.x + v[0]
				newY := curr.y + v[1]

				if newX >= 0 && newX < n && newY >= 0 && newY < m && grid[newX][newY] == '1' && visited[newX][newY] == 0 {
					q.Enqueue(Coord{newX, newY})
					visited[newX][newY] = 1
				}
			}

			q.Dequeue()
		}
	}

	islands := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' && visited[i][j] == 0 {
				bfs(i, j)
				islands++
			}
		}
	}

	return islands
}
