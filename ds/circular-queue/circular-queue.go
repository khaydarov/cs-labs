package circular_queue

type CircularQueue struct {
	data     []int
	head     int
	tail     int
	length   int
	capacity int
}

// EnQueue adds new element to the queue (to the end)
func (q *CircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}

	if q.IsEmpty() {
		q.head = 0
		q.tail = 0
	} else {
		q.tail++
		q.tail %= q.capacity
	}

	q.data[q.tail] = value
	q.length++

	return true
}

// DeQueue removes element from the beginning
func (q *CircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.head++
	q.head %= q.capacity
	q.length--

	return true
}

// Front returns the first element
func (q *CircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.head]
}

// Rear returns the last element
func (q *CircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.tail]
}

// IsEmpty returns bool if queue is empty
func (q *CircularQueue) IsEmpty() bool {
	return q.length == 0
}

// IsFull returns bool if queue is full
func (q *CircularQueue) IsFull() bool {
	return q.length == q.capacity
}

// Construct creates new CircularQueue with capacity of k
func Construct(k int) CircularQueue {
	data := make([]int, k)

	return CircularQueue{
		data,
		-1,
		-1,
		0,
		k,
	}
}
