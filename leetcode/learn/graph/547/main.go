package main

type DisjointSet struct {
	rank   []int
	parent []int
}

func (s *DisjointSet) Find(x int) int {
	if x == s.parent[x] {
		return x
	}

	s.parent[x] = s.Find(s.parent[x])
	return s.parent[x]
}

func (s *DisjointSet) Union(x, y int) {
	xRoot := s.Find(x)
	yRoot := s.Find(y)

	if xRoot == yRoot {
		return
	}

	if s.parent[xRoot] > s.parent[yRoot] {
		s.parent[yRoot] = xRoot
	} else if s.parent[xRoot] < s.parent[yRoot] {
		s.parent[xRoot] = yRoot
	} else {
		s.parent[xRoot] = yRoot
		s.rank[xRoot]++
	}
}

func (s *DisjointSet) Connected(x, y int) bool {
	return s.Find(x) == s.Find(y)
}

func NewDisjointSet(n int) *DisjointSet {
	rank := make([]int, n)
	parent := make([]int, n)

	for i := 0; i < n; i++ {
		rank[i] = 1
		parent[i] = i
	}
	return &DisjointSet{
		rank,
		parent,
	}
}

// TC: O(N^2)
// SC: O(N)
func findCircleNum(isConnected [][]int) int {
	disjointSet := NewDisjointSet(len(isConnected))
	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 1 && i != j {
				disjointSet.Union(i, j)
			}
		}
	}

	var result int
	for i := 0; i < len(isConnected); i++ {
		if disjointSet.parent[i] == i {
			result++
		}
	}
	return result
}

// TC: O(N^2)
// SC: O(N)
func findCircleNum1(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))

	var helper func(x int)
	helper = func(x int) {
		visited[x] = true
		for i := 0; i < len(isConnected); i++ {
			if isConnected[x][i] == 1 && i != x && visited[i] == false {
				helper(i)
			}
		}
	}

	var result int
	for i := 0; i < len(isConnected); i++ {
		if visited[i] {
			continue
		}

		helper(i)
		result++
	}

	return result
}
