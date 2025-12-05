package main

// queue keeps track of the order of cache keys
type queue []int

func (q *queue) push(i int) {
	*q = append(*q, i)
}

func (q *queue) pop() int {
	popValue := (*q)[0]
	*q = (*q)[1:]
	return popValue
}

type FIFOCache struct {
	capacity int

	storage map[int]int
	queue   queue
}

func NewFIFOCache(capacity int) *FIFOCache {
	return &FIFOCache{
		capacity: capacity,
		storage:  make(map[int]int),
		queue:    make(queue, 0),
	}
}

// Get returns value from cache by key
// Returns -1 if key is not in cache
func (c *FIFOCache) Get(key int) int {
	if c.capacity == 0 {
		return -1
	}

	if value, exists := c.storage[key]; exists {
		return value
	}

	return -1
}

// Put adds a key-value pair to the cache
// If the cache is full, the oldest key is removed
func (c *FIFOCache) Put(key, value int) {
	if c.capacity == 0 {
		return
	}

	if _, exists := c.storage[key]; exists {
		c.storage[key] = value
		return
	}

	if len(c.queue) >= c.capacity {
		delete(c.storage, c.queue.pop())
	}

	c.queue.push(key)
	c.storage[key] = value
}
