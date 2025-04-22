package main

// QuickFindDisjointSet is a disjoint set implementation that uses a quick find algorithm.
// It is a simple implementation that is easy to understand and implement.
// However, it has a disadvantage that the union operation is slow because it requires iterating through all elements.
type QuickFindDisjointSet struct {
	parent []int
}

// Find returns the root of the set that x belongs to.
// TC: O(1)
// SC: O(1)
func (s *QuickFindDisjointSet) Find(x int) int {
	return s.parent[x]
}

// Union merges the sets that x and y belong to.
// This implementation is not efficient because it requires iterating through all elements and updating the parent of each element.
// TC: O(n)
// SC: O(1)
func (s *QuickFindDisjointSet) Union(x, y int) {
	xRoot := s.Find(x)
	yRoot := s.Find(y)

	if xRoot != yRoot {
		for i := 0; i < len(s.parent); i++ {
			if s.parent[i] == yRoot {
				s.parent[i] = xRoot
			}
		}
	}
}

// Connected returns true if x and y are in the same set.
// TC: O(1)
// SC: O(1)
func (s *QuickFindDisjointSet) Connected(x, y int) bool {
	return s.Find(x) == s.Find(y)
}
