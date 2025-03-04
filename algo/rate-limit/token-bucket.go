package main

import (
	"sync"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TokenBucket is a rate limiter strategy that uses a token bucket algorithm.
type TokenBucket struct {
	rate       int
	capacity   int
	tokens     int
	lastRefill time.Time

	mu sync.Mutex
}

func NewTokenBucket(rate int, capacity int) RateLimiterStrategy {
	return &TokenBucket{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.refill()
	if tb.tokens < 1 {
		return false
	}

	tb.tokens--
	return true
}

func (tb *TokenBucket) refill() {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	delta := now.UnixMilli() - tb.lastRefill.UnixMilli()

	newTokens := int(delta) / 1000 * tb.rate
	if newTokens > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+newTokens)
		tb.lastRefill = now
	}
}
