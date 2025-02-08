package main

type Queue struct {
	elements []string
}

func (q *Queue) Enqueue(element string) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Dequeue() string {
	if len(q.elements) == 0 {
		return ""
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func NewQueue() *Queue {
	return &Queue{
		elements: make([]string, 0),
	}
}
