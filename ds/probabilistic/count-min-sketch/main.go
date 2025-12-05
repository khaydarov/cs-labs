package count_min_sketch

import (
	"math"

	"github.com/spaolacci/murmur3"
)

type Matrix [][]int

type CountMinSketch struct {
	depth int
	width int

	matrix Matrix
}

func NewCountMinSketch(depth int, width int) *CountMinSketch {
	matrix := make(Matrix, depth)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}

	return &CountMinSketch{
		depth:  depth,
		width:  width,
		matrix: matrix,
	}
}

func (c *CountMinSketch) Insert(item string) {
	for i := 0; i < c.depth; i++ {
		hash := murmur3.New32WithSeed(uint32(i))
		hash.Write([]byte(item))
		digest := hash.Sum32() % uint32(c.width)
		c.matrix[i][digest]++
	}
}

func (c *CountMinSketch) Query(item string) int {
	minCount := math.MaxInt
	for i := 0; i < c.depth; i++ {
		hash := murmur3.New32WithSeed(uint32(i))
		hash.Write([]byte(item))
		digest := hash.Sum32() % uint32(c.width)

		minCount = min(minCount, c.matrix[i][digest])
	}

	return minCount
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
