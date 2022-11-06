package main

import "fmt"

// New creates new LRUCache instance
func New(capacity int) LRUCache {
	head := &item{}
	tail := &item{}

	head.next = tail
	tail.prev = head

	return LRUCache{
		storage:  make(map[int]*item, capacity),
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

// item is doubly-linked list structure
type item struct {
	key   int
	value int

	next *item
	prev *item
}

// LRUCache is a structure that implements LRU caching strategy
type LRUCache struct {
	// key-value storage
	storage map[int]*item

	// References to doubly-linked list items: head nad tail
	head *item
	tail *item

	// Cache capacity and its current size
	capacity int
	size     int
}

// Put adds new element according to the LRU strategy
func (c *LRUCache) Put(key, value int) {
	if item, ok := c.storage[key]; ok {
		item.value = value
		c.bringToMostUsed(key)
	} else if c.capacity == 0 {
		return
	} else {
		if c.isFull() {
			c.removeLeastUsed()
		}
		c.append(key, value)
	}
}

// Get returns value from cache by key
func (c *LRUCache) Get(key int) int {
	if _, ok := c.storage[key]; !ok {
		return -1
	}

	returnValue := c.storage[key]
	c.bringToMostUsed(key)

	return returnValue.value
}

// bringToMostUsed makes key most used
func (c *LRUCache) bringToMostUsed(key int) {
	node := c.storage[key]

	mostUsed := c.tail.prev
	if node == mostUsed {
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = mostUsed
	mostUsed.next = node

	c.tail.prev = node
	node.next = c.tail
}

// removeLeastUsed removes least used key
func (c *LRUCache) removeLeastUsed() {
	delete(c.storage, c.head.next.key)

	leastUsed := c.head.next
	leastUsed.next.prev = c.head
	c.head.next = leastUsed.next

	c.size--
}

// isFull returns bool if length of storage equals to capacity
func (c *LRUCache) isFull() bool {
	return len(c.storage) >= c.capacity
}

// append returns adds to the tail
func (c *LRUCache) append(key, value int) {
	node := &item{key: key, value: value}

	mostUsed := c.tail.prev
	mostUsed.next = node
	node.prev = mostUsed

	c.tail.prev = node
	node.next = c.tail

	c.storage[key] = node
	c.size++
}

func (c *LRUCache) debug() {
	curr := c.head
	for curr != nil {
		fmt.Println(
			fmt.Sprintf("curr.key: %d, curr.val: %d, next: %v, prev: %v", curr.key, curr.value, curr.next, curr.prev),
		)
		curr = curr.next
	}
}
