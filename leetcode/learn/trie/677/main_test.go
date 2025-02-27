package main

import "testing"

func TestSolution(t *testing.T) {
	mapSum := Constructor()
	mapSum.Insert("apple", 3)
	if got := mapSum.Sum("ap"); got != 3 {
		t.Errorf("Test case 0 failed: expected %v but got %v", 3, got)
	}

	mapSum.Insert("app", 2)
	if got := mapSum.Sum("ap"); got != 5 {
		t.Errorf("Test case 1 failed: expected %v but got %v", 5, got)
	}

	mapSum.Insert("apple", 2)
	if got := mapSum.Sum("ap"); got != 4 {
		t.Errorf("Test case 2 failed: expected %v but got %v", 4, got)
	}
}
