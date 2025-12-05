package main

type cacheItem struct {
	key   int
	value int
	freq  int

	prev *cacheItem
	next *cacheItem
}

type LFUCache struct {
	capacity int
	size     int
	items    map[int]*cacheItem
	freqs    map[int]*cacheItem
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		items:    make(map[int]*cacheItem),
		freqs:    make(map[int]*cacheItem),
	}
}

func (c *LFUCache) Get(key int) int {
	if c.capacity == 0 {
		return -1
	}

	if item, exists := c.items[key]; exists {
		c.increment(item)
		return item.value
	}
	return -1
}

func (c *LFUCache) Put(key, value int) {
	if c.capacity == 0 {
		return
	}

	if item, exists := c.items[key]; exists {
		item.value = value
		c.increment(item)
		return
	}

	if c.size >= c.capacity {
		c.removeLeastFreq()
	}

	defaultFreq := 1
	c.items[key] = &cacheItem{
		key:   key,
		value: value,
		freq:  defaultFreq,
	}
	head := c.freqs[defaultFreq]
	if head == nil {
		c.freqs[defaultFreq] = c.items[key]
	} else {
		head.next = c.items[key]
		c.items[key].prev = head
	}
	c.size++
}

func (c *LFUCache) increment(item *cacheItem) {
	head := c.freqs[item.freq]
	if head == item {
		c.freqs[item.freq] = item.next
		if item.next != nil {
			item.next.prev = nil
		}
	} else {
		if item.next != nil {
			item.next.prev = item.prev
		}
		if item.prev != nil {
			item.prev.next = item.next
		}
	}

	if c.freqs[item.freq] == nil {
		delete(c.freqs, item.freq)
	}

	nextFreq := item.freq + 1
	if c.freqs[nextFreq] == nil {
		c.freqs[nextFreq] = &cacheItem{
			key:   item.key,
			value: item.value,
			freq:  nextFreq,
		}
	} else {
		current := c.freqs[nextFreq]
		for current.next != nil {
			current = current.next
		}
		current.next = item
		item.prev = current
	}
}

func (c *LFUCache) removeLeastFreq() {
	var leastFreq *cacheItem
	for i := 1; i < len(c.freqs); i++ {
		if c.freqs[i] != nil {
			leastFreq = c.freqs[i]
			break
		}
	}

	if leastFreq == nil {
		return
	}

	delete(c.items, leastFreq.key)
	if leastFreq.next == nil {
		delete(c.freqs, leastFreq.freq)
	} else {
		c.freqs[leastFreq.freq] = leastFreq.next
		leastFreq.next.prev = nil
	}
	c.size--
}
