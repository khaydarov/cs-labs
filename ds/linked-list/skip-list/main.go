package main

import (
	"math/rand"
)

type Node struct {
	Val  int
	Next []*Node
}

func newNode(val, level int) *Node {
	return &Node{
		Val:  val,
		Next: make([]*Node, level),
	}
}

type SkipList struct {
	head     *Node
	level    int
	maxLevel int
}

func NewSkipList(maxLevel int) *SkipList {
	return &SkipList{
		head:     newNode(0, maxLevel),
		level:    1,
		maxLevel: maxLevel,
	}
}

func (sl *SkipList) Add(val int) {
	update := make([]*Node, sl.maxLevel)
	current := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for current.Next[i] != nil && current.Next[i].Val < val {
			current = current.Next[i]
		}
		update[i] = current
	}

	if next := update[0].Next[0]; next != nil && next.Val == val {
		return
	}

	level := sl.randomLvl()

	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.head
		}
		sl.level = level
	}

	n := newNode(val, level)
	for i := 0; i < level; i++ {
		n.Next[i] = update[i].Next[i]
		update[i].Next[i] = n
	}
}

func (sl *SkipList) Search(val int) bool {
	current := sl.head
	for i := len(current.Next) - 1; i >= 0; i-- {
		for current.Next[i] != nil && current.Next[i].Val < val {
			current = current.Next[i]
		}
	}

	return current.Next[0] != nil && current.Next[0].Val == val
}

func (sl *SkipList) Delete(val int) {
	current := sl.head
	update := make([]*Node, sl.maxLevel)

	for i := sl.level - 1; i >= 0; i-- {
		for current.Next[i] != nil && current.Next[i].Val < val {
			current = current.Next[i]
		}
		update[i] = current
	}

	target := update[0].Next[0]
	if target == nil || target.Val != val {
		return
	}

	for i := 0; i < len(target.Next); i++ {
		if update[i].Next[i] == target {
			update[i].Next[i] = target.Next[i]
		}
	}

	for sl.level > 1 && sl.head.Next[sl.level-1] == nil {
		sl.level--
	}
}

func (sl *SkipList) randomLvl() int {
	lvl := 1
	for rand.Float64() < 0.5 && lvl < sl.maxLevel {
		lvl++
	}

	return lvl
}
