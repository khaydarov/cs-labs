package main

// Set structure
type Set struct {
	data   []int
	length int
}

// Add Adds new element to the set
// Time complexity: O(N)
// Note: We can use balanced tree to search with O(Log N)
func (s *Set) Add(x int) bool {
	if s.Has(x) {
		return false
	}

	s.data = append(s.data, x)
	s.length++

	return true
}

// Has Checks if Set contains element
func (s *Set) Has(x int) bool {
	for _, v := range s.data {
		if v == x {
			return true
		}
	}

	return false
}

func NewSet() *Set {
	return &Set{}
}
