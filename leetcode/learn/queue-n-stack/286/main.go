package main

import "math"

type Coord [2]int

type Queue struct {
	coordinates []*Coord
}

func (s *Queue) Enqueu(x, y int) {
	coord := &Coord{x, y}

	s.coordinates = append(s.coordinates, coord)
}

func (s *Queue) Dequeue() *Coord {
	if s.Empty() {
		return nil
	}

	front := s.coordinates[0]
	s.coordinates = s.coordinates[1:]

	return front
}

func (s *Queue) Empty() bool {
	return len(s.coordinates) == 0
}

func (s *Queue) Size() int {
	return len(s.coordinates)
}

func wallsAndGates(rooms [][]int) {
	var coordinates []*Coord
	q := Queue{
		coordinates,
	}

	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				q.Enqueu(i, j)
			}
		}
	}

	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for !q.Empty() {
		point := q.Dequeue()
		row := point[0]
		col := point[1]

		for _, dir := range dirs {
			r := row + dir[0]
			c := col + dir[1]

			if r < 0 || c < 0 || r >= len(rooms) || c >= len(rooms[0]) || rooms[r][c] != math.MaxInt32 {
				continue
			}

			rooms[r][c] = rooms[row][col] + 1
			q.Enqueu(r, c)
		}
	}
}
