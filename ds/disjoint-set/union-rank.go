package main

// UnionRankDisjointSet is a disjoint set implementation that uses a union by rank algorithm.
// It is a simple implementation that is easy to understand and implement.
// However, it has a disadvantage that the union operation is slow because it requires iterating through all elements.
type UnionRankDisjointSet struct {
	parent []int
	rank   []int
}

// Find returns the root of the set that x belongs to.
// TC: O(log n)
// SC: O(1)
func (s *UnionRankDisjointSet) Find(x int) int {
	if x == s.parent[x] {
		return x
	}
	s.parent[x] = s.Find(s.parent[x])
	return s.parent[x]
}

// Union merges the sets that x and y belong to.
// TC: O(log n)
// SC: O(1)
func (s *UnionRankDisjointSet) Union(x, y int) {
	xRoot := s.Find(x)
	yRoot := s.Find(y)

	if xRoot == yRoot {
		return
	}

	if s.rank[xRoot] > s.rank[yRoot] {
		s.parent[yRoot] = xRoot
	} else if s.rank[xRoot] < s.rank[yRoot] {
		s.parent[xRoot] = yRoot
	} else {
		s.parent[yRoot] = xRoot
		s.rank[xRoot]++
	}
}

// Connected returns true if x and y are in the same set.
// TC: O(log n)
// SC: O(1)
func (s *UnionRankDisjointSet) Connected(x, y int) bool {
	return s.Find(x) == s.Find(y)
}
