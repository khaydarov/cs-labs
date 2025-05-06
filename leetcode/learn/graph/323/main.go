package main

import "fmt"

type DisjointSet struct {
	parent []int
	rank   []int
	count  int
}

func NewDisjointSet(n int) *DisjointSet {
	rank := make([]int, n)
	parent := make([]int, n)

	for i := 0; i < n; i++ {
		rank[i] = 1
		parent[i] = i
	}

	return &DisjointSet{
		parent: parent,
		rank:   rank,
		count:  n,
	}
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

	if s.rank[xRoot] > s.rank[yRoot] {
		s.parent[yRoot] = xRoot
	} else if s.rank[xRoot] < s.rank[yRoot] {
		s.parent[xRoot] = yRoot
	} else {
		s.parent[xRoot] = yRoot
		s.rank[xRoot]++
	}
	s.count--
}

func countComponents1(n int, edges [][]int) int {
	disjointSet := NewDisjointSet(n)
	for _, edge := range edges {
		disjointSet.Union(edge[0], edge[1])
	}

	return disjointSet.count
}

func countComponents(n int, edges [][]int) int {
	visited := make([]bool, n)
	var helper func(x int)
	helper = func(x int) {
		visited[x] = true

		for i := 0; i < len(edges); i++ {
			if edges[i][0] == x && !visited[edges[i][1]] {
				helper(edges[i][1])
			} else if edges[i][1] == x && !visited[edges[i][0]] {
				helper(edges[i][0])
			}
		}
	}

	var result int
	for i := 0; i < n; i++ {
		if !visited[i] {
			helper(i)
			result++
		}
	}

	return result
}

func main() {
	edges := [][]int{
		//{0, 1},
		//{1, 2},
		//{2, 3},
		//{3, 4},
	}
	fmt.Println(countComponents(5, edges))
}
