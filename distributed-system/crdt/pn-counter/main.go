package main

import "fmt"

type PNCounter struct {
	positive *GCounter
	negative *GCounter
}

func NewPNCounter() *PNCounter {
	return &PNCounter{
		positive: NewGCounter(),
		negative: NewGCounter(),
	}
}

func (p *PNCounter) Increment(node string) {
	p.positive.Increment(node)
}

func (p *PNCounter) Decrement(node string) {
	p.negative.Increment(node)
}

func (p *PNCounter) Merge(other *PNCounter) {
	p.positive.Merge(other.positive)
	p.negative.Merge(other.negative)
}

func (p *PNCounter) Value() int {
	return p.positive.Value() - p.negative.Value()
}

func main() {
	counter := NewPNCounter()
	counter.Increment("A")
	counter.Increment("B")
	counter.Decrement("A")
	counter.Decrement("A")

	fmt.Println("Значение счетчика:", counter.Value())
}
