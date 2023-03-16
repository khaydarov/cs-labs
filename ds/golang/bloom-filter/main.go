package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	size        int
	probability float64
	hashCount   int
	bitArray    []bool
}

func NewBloomFilter(numberOfItems int, probability float64) *BloomFilter {
	size := calculateBFSize(numberOfItems, probability)
	bitArray := make([]bool, size)

	return &BloomFilter{
		size:        size,
		probability: probability,
		hashCount:   calculateHashCount(numberOfItems, size),
		bitArray:    bitArray,
	}
}

func (b *BloomFilter) Insert(item string) {
	for i := 0; i < b.hashCount; i++ {
		hasher := murmur3.New32WithSeed(uint32(i))
		hasher.Write([]byte(item))

		digest := hasher.Sum32() % uint32(b.size)
		b.bitArray[digest] = true
	}
}

func (b *BloomFilter) Check(item string) bool {
	for i := 0; i < b.hashCount; i++ {
		hasher := murmur3.New32WithSeed(uint32(i))
		hasher.Write([]byte(item))

		digest := hasher.Sum32() % uint32(b.size)
		if b.bitArray[digest] == false {
			return false
		}
	}

	return true
}

func main() {
	bloomFilter := NewBloomFilter(10, 0.05)
	bloomFilter.Insert("hello")
	bloomFilter.Insert("Hi")

	fmt.Println(bloomFilter.Check("hello"))
}
