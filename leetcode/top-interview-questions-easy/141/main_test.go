package main

import "testing"

func buildCycleList() *ListNode {
	one := &ListNode{3, nil}
	two := &ListNode{2, nil}
	three := &ListNode{2, nil}
	four := &ListNode{5, nil}
	//five := &ListNode{3, nil}

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = one

	return one
}

func TestSolution(t *testing.T) {
	list := buildCycleList()
	got := hasCycle(list)
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
