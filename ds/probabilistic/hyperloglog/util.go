package main

import (
	"github.com/spaolacci/murmur3"
)

func hash(item string) uint32 {
	hash := murmur3.New32()
	hash.Write([]byte(item))
	return hash.Sum32()
}

func leadingZeroCount(x uint32) uint8 {
	var count uint8 = 1
	for (x & 1) == 0 {
		count++
		x >>= 1
	}
	return count
}
