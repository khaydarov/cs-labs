package main

import (
	"math"
)

type HyperLogLog struct {
	p         int
	m         int
	registers []uint8
}

func NewHyperLogLog(p int) *HyperLogLog {
	return &HyperLogLog{
		registers: make([]uint8, 1<<p),
		p:         p,
		m:         1 << p,
	}
}

func (h *HyperLogLog) Add(item string) {
	hash := hash(item)
	registerIndex := hash >> (32 - h.p)

	w := hash << h.p
	w = w >> h.p
	leftmost1Position := leadingZeroCount(w) + 1
	if leftmost1Position > h.registers[registerIndex] {
		h.registers[registerIndex] = leftmost1Position
	}
}

func (h *HyperLogLog) Count() uint64 {
	sum := 0.0
	for _, reg := range h.registers {
		sum += math.Pow(2.0, -float64(reg))
	}

	alpha := 0.7213 / (1.0 + 1.079/float64(h.m))
	estimate := alpha * float64(h.m*h.m) / sum

	if estimate <= 2.5*float64(h.m) {
		zeros := 0
		for _, reg := range h.registers {
			if reg == 0 {
				zeros++
			}
		}
		if zeros > 0 {
			estimate = float64(h.m) * math.Log(float64(h.m)/float64(zeros))
		}
	}

	return uint64(estimate + 0.5)
}
