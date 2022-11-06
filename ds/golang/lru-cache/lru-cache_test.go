package main

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	// Test case
	//lru := New(2)
	//lru.Put(1,1)
	//lru.Put(2,2)
	//fmt.Println(lru.Get(1))
	//lru.Put(3,3)
	//fmt.Println(lru.Get(2))
	//lru.Put(4, 4)
	//fmt.Println(lru.Get(1))
	//fmt.Println(lru.Get(3))
	//fmt.Println(lru.Get(4))

	// Test case
	//lru := New(2)
	//fmt.Println(lru.Get(2))
	//lru.Put(2, 6)
	//fmt.Println(lru.Get(1))
	//lru.Put(1, 5)
	//lru.Put(1, 2)
	//fmt.Println(lru.Get(1))
	//fmt.Println(lru.Get(2))

	// Test case
	//lru := New(1)
	//lru.Put(2, 1)
	//fmt.Println(lru.Get(2))
	//lru.debug()
	//lru.Put(3, 2)
	//fmt.Println(lru.Get(2))
	//fmt.Println(lru.Get(3))

	// Test case
	lru := New(10)
	lru.Put(10, 13)
	lru.Put(3, 17)
	lru.Put(6, 11)
	lru.Put(10, 5)
	lru.Put(9, 10)
	fmt.Println(lru.Get(13))
	lru.Put(2, 19)
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(3))
	lru.Put(5, 25)
	fmt.Println(lru.Get(8))
	lru.Put(9, 22)
	lru.Put(5, 5)
	lru.Put(1, 30)
	fmt.Println(lru.Get(11))
	lru.Put(9, 12)
	fmt.Println(lru.Get(7))
	fmt.Println(lru.Get(5))
	fmt.Println(lru.Get(8))
	fmt.Println(lru.Get(9))
	lru.Put(12, 24)
	fmt.Println(lru.size)

	// Test case
	//lru := New(1)
	//fmt.Println(lru.Get(6))
	//fmt.Println(lru.Get(8))
	//lru.Put(12, 1)
	//fmt.Println(lru.Get(2))
	//lru.Put(15, 11)
	//lru.Put(5, 2)
	//lru.Put(1, 15)
	//lru.Put(4, 2)
	////lru.debug()
	//fmt.Println(lru.Get(5))
	//lru.Put(15, 15)
}
