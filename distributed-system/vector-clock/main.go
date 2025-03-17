package main

type Node string

const (
	Less       = -1
	Greater    = 1
	Equal      = 0
	Concurrent = 2
)

type VectorClock map[Node]int

func NewVectorClock() VectorClock {
	return make(VectorClock)
}

func (vc VectorClock) Increment(node Node) {
	vc[node]++
}

func (vc VectorClock) Merge(other VectorClock) {
	for node, clock := range other {
		if current, exists := vc[node]; !exists || clock > current {
			vc[node] = clock
		}
	}
}

func (vc VectorClock) Compare(other VectorClock) int {
	allNodes := make(map[Node]bool)
	for node := range vc {
		allNodes[node] = true
	}

	for node := range other {
		allNodes[node] = true
	}

	vLess, vGreater := false, false
	for node := range allNodes {
		currentTime := vc[node]
		otherTime := other[node]

		if currentTime < otherTime {
			vLess = true
		}

		if currentTime > otherTime {
			vGreater = true
		}
	}

	switch {
	case !vLess && !vGreater:
		return Equal
	case vLess && !vGreater:
		return Less
	case vGreater && !vLess:
		return Greater
	default:
		return Concurrent
	}
}
