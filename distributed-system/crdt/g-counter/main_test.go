package main

import "testing"

func TestGCounter(t *testing.T) {
	counter := NewGCounter()
	counter.Increment("A")
	counter.Increment("B")
	counter.Increment("A")

	if counter.Value() != 3 {
		t.Errorf("Counter should be 3")
	}
}

func TestGCounterMerge(t *testing.T) {
	counter1 := NewGCounter()
	counter2 := NewGCounter()

	counter1.Increment("A")
	counter1.Increment("A")
	counter2.Increment("B")

	counter1.Merge(counter2)

	if counter1.Value() != 3 {
		t.Errorf("Counter1 should be 3")
	}

	if counter2.Value() != 1 {
		t.Errorf("Counter2 should be 1")
	}
}
