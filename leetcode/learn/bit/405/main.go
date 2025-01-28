package main

import (
	"strings"
)

type Builder struct {
	strings.Builder
}

func (b *Builder) Reverse() {
	runes := []rune(b.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	b.Reset()
	b.WriteString(string(runes))
}

func toHexValue(v int) byte {
	if v < 10 {
		return byte(v + '0')
	}
	return byte(v - 10 + 'a')
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	var b Builder
	for i := 0; i < 8; i++ {
		b.WriteByte(toHexValue(num & 0xf))
		num >>= 4
	}

	b.Reverse()
	result := b.String()
	i := 0
	for result[i] == '0' {
		i++
	}

	return result[i:]
}
