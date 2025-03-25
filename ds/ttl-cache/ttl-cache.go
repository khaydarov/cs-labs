package main

import (
	"sync"
	"time"
)

type item struct {
	value      interface{}
	expiration time.Time
}

type ttlCache struct {
	items map[string]*item
	mu    sync.RWMutex
}

func NewTTLCache() *ttlCache {
	return &ttlCache{
		items: make(map[string]*item),
	}
}

func (c *ttlCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.items[key]; ok {
		c.items[key].expiration = time.Now().Add(ttl)
		c.items[key].value = value
		return
	}

	c.items[key] = &item{
		value:      value,
		expiration: time.Now().Add(ttl),
	}
}

func (c *ttlCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	if time.Now().After(item.expiration) {
		delete(c.items, key)
		return nil, false
	}

	return item.value, true
}

func (c *ttlCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}
