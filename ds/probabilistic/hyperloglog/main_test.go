package main

import (
	"testing"
)

func TestHyperLogLog(t *testing.T) {
	hll := NewHyperLogLog(16)
	hll.Add("hello")
	hll.Add("hello")
	hll.Add("world")
	if hll.Count() != 2 {
		t.Errorf("Expected count to be 2, got %d", hll.Count())
	}

	hll.Add("hello")
	hll.Add("hello")
	hll.Add("world")
	if hll.Count() != 2 {
		t.Errorf("Expected count to be 2, got %d", hll.Count())
	}

	hll.Add("bye!")
	if hll.Count() != 3 {
		t.Errorf("Expected count to be 3, got %d", hll.Count())
	}
}
