package main

import (
	"testing"
)

func TestOrSet(t *testing.T) {
	orSet := NewOrSet()

	orSet.Add("A")
	orSet.Add("B")
	orSet.Add("C")

	if !orSet.Contains("A") {
		t.Errorf("A should be in the set")
	}

	if !orSet.Contains("B") {
		t.Errorf("B should be in the set")
	}

	orSet.Remove("A")
	if orSet.Contains("A") {
		t.Errorf("A should be removed from the set")
	}

	if orSet.Size() != 2 {
		t.Errorf("Size should be 2")
	}
}

func TestOrSetMerge(t *testing.T) {
	orSet1 := NewOrSet()
	orSet2 := NewOrSet()

	orSet1.Add("A")
	orSet1.Add("B")

	orSet2.Add("B")
	orSet2.Add("C")

	orSet1.Merge(orSet2)

	if orSet1.Size() != 3 {
		t.Errorf("Size should be 3")
	}

	if !orSet1.Contains("C") {
		t.Errorf("C should be in the set")
	}
}
