package circular_queue

import (
	"reflect"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	expect := []interface{}{true, 6, 6, true, true, 5, true, -1, false, false, false}
	var output []interface{}
	q := Construct(6)
	output = append(output, q.EnQueue(6))
	output = append(output, q.Rear())
	output = append(output, q.Rear())
	output = append(output, q.DeQueue())
	output = append(output, q.EnQueue(5))
	output = append(output, q.Rear())
	output = append(output, q.DeQueue())
	output = append(output, q.Front())
	output = append(output, q.DeQueue())
	output = append(output, q.DeQueue())
	output = append(output, q.DeQueue())

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("Wrong answer")
	}
}
