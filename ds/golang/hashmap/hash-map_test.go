package hashmap

import (
	"testing"
)

func TestHashMap(t *testing.T) {
	h := Construct(10)
	h.Put(1, 1)
	h.Put(2, 2)
	h.Put(3, 3)

	if h.Get(1) != 1 {
		t.Errorf("Key `1` must have value of 1")
	}

	if h.Get(2) != 2 {
		t.Errorf("Key `2` must have value of 2")
	}

	if h.Get(3) != 3 {
		t.Errorf("Key `3` must have value of 3")
	}
}
