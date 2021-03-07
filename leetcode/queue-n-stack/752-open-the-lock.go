package main

import "fmt"

type WheelPosition struct {
	x, y, z, w int
}

type Queue struct {
	data []WheelPosition
}

func (q *Queue) Enqueue(c WheelPosition) {
	q.data = append(q.data, c)
}

func (q *Queue) Dequeue() {
	q.data = q.data[1:]
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue) Front() WheelPosition {
	return q.data[0]
}

func openLock(deadends []string, target string) int {
	locks := make(map[string]bool)
	for _, lock := range deadends {
		locks[lock] = true
	}

	if _, ok := locks["0000"]; ok {
		return -1
	}

	if _, ok := locks[target]; ok {
		return -1
	}

	dirs := [][]int{
		{0, 0, 0, 1},
		{0, 0, 1, 0},
		{0, 1, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 0, -1},
		{0, 0, -1, 0},
		{0, -1, 0, 0},
		{-1, 0, 0, 0},
	}

	var code string
	visited := make(map[string]bool)

	steps := 0
	q := Queue{}
	q.Enqueue(WheelPosition{0, 0, 0, 0})

	for !q.Empty() {
		l := len(q.data)

		for i := 0; i < l; i++ {
			current := q.Front()
			code = fmt.Sprintf("%d%d%d%d", current.x, current.y, current.z, current.w)

			if code == target {
				return steps
			}

			for _, dir := range dirs {
				newX := current.x + dir[0]
				newY := current.y + dir[1]
				newZ := current.z + dir[2]
				newW := current.w + dir[3]

				if newX > 9 {
					newX = 0
				} else if newX == -1 {
					newX = 9
				}

				if newY > 9 {
					newY = 0
				} else if newY == -1 {
					newY = 9
				}

				if newZ > 9 {
					newZ = 0
				} else if newZ == -1 {
					newZ = 9
				}

				if newW > 9 {
					newW = 0
				} else if newW == -1 {
					newW = 9
				}

				code = fmt.Sprintf("%d%d%d%d", newX, newY, newZ, newW)

				if v, ok := visited[code]; !ok && v == false {
					visited[code] = true

					if _, ok := locks[code]; !ok {
						q.Enqueue(WheelPosition{newX, newY, newZ, newW})
					}
				}
			}
			q.Dequeue()
		}
		steps++
	}

	return -1
}

func main() {
	d := []string{
		//"0201","0101","0102","1212","2002",
		//"8888",
		//"8887","8889","8878","8898","8788","8988","7888","9888",
		"0000",
	}
	r := openLock(d, "8888")
	fmt.Println(r)
}