package main

// Set structure
// Time complexity depends on DS of storage. Current implementation uses simple array
// For optimization we can use `balanced tree`
type Set struct {
	data   []int
	length int
}

// Add appends new element to the set
// Time complexity: O(N)
func (s *Set) Add(x int) bool {
	if s.Has(x) {
		return false
	}

	s.data = append(s.data, x)
	s.length++

	return true
}

// Has checks if Set contains element
func (s *Set) Has(x int) bool {
	for _, v := range s.data {
		if v == x {
			return true
		}
	}

	return false
}

// Values returns Sets' data
func (s *Set) Values() []int {
	return s.data
}

// Union merges two Sets
// Time complexity: O(max(N, M)), where N, M are Sets size
func (s *Set) Union(anotherSet Set) Set {
	unionSet := Set{}

	for _, v := range s.Values() {
		unionSet.Add(v)
	}

	for _, v := range anotherSet.Values() {
		unionSet.Add(v)
	}

	return unionSet
}

// Intersection returns intersection between two Sets
// Time complexity: O(N*M)
func (s *Set) Intersection(anotherSet Set) Set {
	intersectionSet := Set{}
	for _, v := range s.Values() {
		if anotherSet.Has(v) {
			intersectionSet.Add(v)
		}
	}

	return intersectionSet
}

// Difference returns difference between two Sets
// Time complexity: O(N*M)
func (s *Set) Difference(anotherSet Set) Set {
	differenceSet := Set{}
	for _, v := range s.Values() {
		if !anotherSet.Has(v) {
			differenceSet.Add(v)
		}
	}

	return differenceSet
}
