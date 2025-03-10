package hashset

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	h := Construct(10)
	h.Add(1)
	h.Add(2)
	h.Add(3)

	if h.Contains(3) == false {
		t.Errorf("Hashset must containt 3")
	}

	if h.Contains(4) == true {
		t.Errorf("Hashset must not contain 4")
	}
}
