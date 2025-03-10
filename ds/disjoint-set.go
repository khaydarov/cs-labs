package main

import (
	"fmt"
)

// DisjointSet Data structure to administrate a set of disjoint sets which is subset of a global set
type DisjointSet struct {
	parent 	[]int
	rank	[]int
}

// Find lookups parent of x
//
// E.g.
// parent: [1, 2, 3, 4, 5]
// x = 3 - parent is 3
//
// parent: [1, 2, 2, 4, 5]
// x = 3 - parent is 2
func (s *DisjointSet) Find(x int) int {
	if s.parent[x] != x {
		return s.Find(s.parent[x])
	}

	return s.parent[x]
}

// Union joins two sets
func (s *DisjointSet) Union(x, y int) {
	xSet := s.Find(x)
	ySet := s.Find(y)

	if xSet == ySet {
		return
	}

	if s.rank[xSet] < s.rank[ySet] {
		s.parent[xSet] = ySet
	} else if s.rank[ySet] > s.rank[xSet] {
		s.parent[ySet] = xSet
	} else {
		s.parent[ySet] = xSet
		s.rank[xSet]++
	}
}

func Construct(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &DisjointSet{parent, rank}
}

func main() {
	disjointSet := Construct(5)
	disjointSet.Union(2, 1)
	disjointSet.Union(2, 3)
	fmt.Println(disjointSet, disjointSet.Find(2))
}