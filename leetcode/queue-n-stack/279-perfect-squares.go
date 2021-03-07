package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(c int) {
	q.data = append(q.data, c)
}

func (q *Queue) Dequeue() {
	q.data = q.data[1:]
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue) Front() int {
	return q.data[0]
}

func (q *Queue) Length() int {
	return len(q.data)
}

func numSquares(n int) int {
	q := Queue{}
	q.Enqueue(n)

	result := 0
	for !q.Empty() {
		l := q.Length()

		identicals := make(map[int]bool)
		for i := 0; i < l; i++ {
			current := q.Front()

			if current == 0 {
				return result
			}

			for j := 1; j * j <= current; j++ {
				if _, ok := identicals[current - j * j]; !ok {
					q.Enqueue(current - j * j)
					identicals[current - j * j] = true
				}

			}

			q.Dequeue()
		}

		result++
	}

	return -1
}

func main() {
	r := numSquares(5)
	fmt.Println(r)
}