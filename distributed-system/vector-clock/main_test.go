package main

import (
	"testing"
)

func TestVectorClock(t *testing.T) {
	clocks := map[string]VectorClock{
		"A": NewVectorClock(),
		"B": NewVectorClock(),
		"C": NewVectorClock(),
		"D": NewVectorClock(),
	}

	clocks["A"].Increment("A")
	clocks["B"].Increment("B")

	if clocks["A"].Compare(clocks["B"]) != Concurrent {
		t.Errorf("A and B should be concurrent")
	}

	clocks["A"].Increment("A")
	clocks["C"].Merge(clocks["A"])
	clocks["C"].Increment("C")

	if clocks["A"].Compare(clocks["C"]) != Less {
		t.Errorf("A and C should be less")
	}
}
