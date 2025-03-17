package main

type GCounter struct {
	counter map[string]int
}

func NewGCounter() *GCounter {
	return &GCounter{counter: make(map[string]int)}
}

func (g *GCounter) Increment(node string) {
	g.counter[node]++
}

func (g *GCounter) Merge(other *GCounter) {
	for node, count := range other.counter {
		if current, exists := g.counter[node]; !exists || count > current {
			g.counter[node] = count
		}
	}
}

func (g *GCounter) Value() int {
	total := 0
	for _, count := range g.counter {
		total += count
	}
	return total
}
