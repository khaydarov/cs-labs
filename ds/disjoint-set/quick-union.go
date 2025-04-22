package main

// QuickUnionDisjointSet is a disjoint set implementation that uses a quick union algorithm.
// It is a simple implementation that is easy to understand and implement.
// However, it has a disadvantage that the union operation is slow because it requires iterating through all elements.
type QuickUnionDisjointSet struct {
	parent []int
}

// Find returns the root of the set that x belongs to.
// TC: O(n)
// SC: O(1)
func (s *QuickUnionDisjointSet) Find(x int) int {
	for s.parent[x] != x {
		x = s.parent[x]
	}
	return x
}

// Union merges the sets that x and y belong to.
// TC: O(n)
// SC: O(1)
func (s *QuickUnionDisjointSet) Union(x, y int) {
	xRoot := s.Find(x)
	yRoot := s.Find(y)

	if xRoot != yRoot {
		s.parent[xRoot] = yRoot
	}
}

// Connected returns true if x and y are in the same set.
// TC: O(n)
// SC: O(1)
func (s *QuickUnionDisjointSet) Connected(x, y int) bool {
	return s.Find(x) == s.Find(y)
}
