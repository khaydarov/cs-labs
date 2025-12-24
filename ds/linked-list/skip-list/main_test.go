package main

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	sl := NewSkipList(4)
	sl.Add(1)
	sl.Add(2)
	sl.Add(3)

	if sl.Search(2) == false {
		t.Errorf("Search 2 failed")
	}

	if sl.Search(5) == true {
		t.Errorf("Search 3 failed")
	}
}
