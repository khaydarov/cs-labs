package main

import "testing"

func TestQuickFindDisjointSet(t *testing.T) {
	disjointSet := Construct(5, QuickFind)
	disjointSet.Union(1, 2)
	disjointSet.Union(2, 3)

	if disjointSet.Find(1) != 1 {
		t.Errorf("Find(1) = %d; want 1", disjointSet.Find(1))
	}
	if disjointSet.Find(2) != 1 {
		t.Errorf("Find(2) = %d; want 1", disjointSet.Find(2))
	}

	if disjointSet.Find(3) != 1 {
		t.Errorf("Find(3) = %d; want 1", disjointSet.Find(3))
	}

	if disjointSet.Find(4) != 4 {
		t.Errorf("Find(4) = %d; want 4", disjointSet.Find(4))
	}
}

func TestQuickUnionDisjointSet(t *testing.T) {
	disjointSet := Construct(5, QuickUnion)
	disjointSet.Union(1, 2)
	disjointSet.Union(2, 3)

	if !disjointSet.Connected(1, 3) {
		t.Errorf("Connected(1, 3) = false; want true")
	}

	if disjointSet.Connected(1, 4) {
		t.Errorf("Connected(1, 4) = true; want false")
	}
}

func TestUnionRankDisjointSet(t *testing.T) {
	disjointSet := Construct(5, UnionRank)
	disjointSet.Union(1, 2)
	disjointSet.Union(2, 3)

	if !disjointSet.Connected(1, 3) {
		t.Errorf("Connected(1, 3) = false; want true")
	}

	if disjointSet.Connected(1, 4) {
		t.Errorf("Connected(1, 4) = true; want false")
	}
}
