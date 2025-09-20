package main

import (
	"crypto/sha256"
)

type Sha256Hasher struct{}

func (s *Sha256Hasher) Hash(input []byte) HashValue {
	sum := sha256.Sum256(input)

	return sum[:]
}

func (s *Sha256Hasher) HashPair(left, right []byte) HashValue {
	buf := make([]byte, 0, len(left)+len(right))
	buf = append(buf, left...)
	buf = append(buf, right...)

	return s.Hash(buf)
}
