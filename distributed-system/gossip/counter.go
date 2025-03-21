package main

type GCounter struct {
	Counter map[NodeId]int `json:"counter"`
}

func NewGCounter() *GCounter {
	return &GCounter{Counter: make(map[NodeId]int)}
}

func (g *GCounter) Increment(nodeId NodeId) {
	g.Counter[nodeId]++
}

func (g *GCounter) Merge(other *GCounter) {
	for node, count := range other.Counter {
		if current, exists := g.Counter[node]; !exists || count > current {
			g.Counter[node] = count
		}
	}
}

func (g *GCounter) Value() int {
	total := 0
	for _, count := range g.Counter {
		total += count
	}
	return total
}
