package main

import "fmt"

type DisjointSet interface {
	Find(x int) int
	Union(x, y int)
	Connected(x, y int) bool
}

const (
	QuickFind = iota
	QuickUnion
	UnionRank
)

func Construct(n int, strategy int) DisjointSet {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	switch strategy {
	case QuickFind:
		return &QuickFindDisjointSet{parent}
	case QuickUnion:
		return &QuickUnionDisjointSet{parent}
	case UnionRank:
		rank := make([]int, n)
		for i := 0; i < n; i++ {
			rank[i] = 1
		}
		return &UnionRankDisjointSet{parent, rank}
	default:
		return nil
	}
}

func main() {
	disjointSet := Construct(5, QuickFind)
	disjointSet.Union(1, 2)
	disjointSet.Union(2, 3)

	fmt.Println(disjointSet)
	fmt.Println(disjointSet.Find(1))
	fmt.Println(disjointSet.Find(2))
	fmt.Println(disjointSet.Find(0))
}
