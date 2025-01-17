package main

import "testing"

func TestSolution(t *testing.T) {
	c := Constructor()
	c.AddNum(1)
	if c.FindMedian() != 1 {
		t.Error("Incorrect median for [1]")
	}

	c.AddNum(2)
	if c.FindMedian() != 1.5 {
		t.Error("Incorrect median for [1, 2]")
	}

	c.AddNum(3)
	if c.FindMedian() != 2 {
		t.Error("Incorrect median for [1, 2, 3]")
	}

	c.AddNum(4)
	if c.FindMedian() != 2.5 {
		t.Errorf("Incorrect median for [1, 2, 3, 4]")
	}
}
