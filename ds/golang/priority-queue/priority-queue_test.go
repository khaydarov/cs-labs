package priority_queue

import (
	"testing"
)

func TestArrayPriorityQueue(t *testing.T) {
	q := ConstructArrayPriorityQueue(5)
	q.Insert(2, 2)
	q.Insert(1, 1)
	q.Insert(9, 11)
}
