package main

import "errors"

type Queue []string

func NewQueue(vals []string) *Queue {
	q := Queue{}
	for _, val := range vals {
		q.Enqueue(val)
	}

	return &q
}

func (q *Queue) Enqueue(val string) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue() (string, error) {
	if len(*q) == 0 {
		return "", errors.New("queue is empty")
	}

	val := (*q)[0]
	*q = (*q)[1:]
	return val, nil
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
