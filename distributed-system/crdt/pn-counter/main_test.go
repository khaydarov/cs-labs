package main

import (
	"testing"
)

func TestPNCounter(t *testing.T) {
	counter := NewPNCounter()
	counter.Increment("A")
	counter.Increment("B")
	counter.Decrement("A")
	counter.Decrement("A")

	if counter.Value() != 0 {
		t.Errorf("Counter should be 0")
	}

	counter.Increment("A")
	counter.Increment("B")
	counter.Increment("C")

	if counter.Value() != 3 {
		t.Errorf("Counter should be 3")
	}

	counter.Decrement("C")
	if counter.Value() != 2 {
		t.Errorf("Counter should be 2")
	}
}

func TestPNCounterMerge(t *testing.T) {
	counter1 := NewPNCounter()
	counter2 := NewPNCounter()

	counter1.Increment("A")
	counter2.Increment("B")

	counter1.Merge(counter2)

	if counter1.Value() != 2 {
		t.Errorf("Counter1 should be 2")
	}

	if counter2.Value() != 1 {
		t.Errorf("Counter2 should be 1")
	}
}
