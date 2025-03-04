package main

import "time"

type RateLimiter struct {
	strategy RateLimiterStrategy
}

type RateLimiterStrategy interface {
	Allow() bool
}

func NewRateLimiter(strategy RateLimiterStrategy) *RateLimiter {
	return &RateLimiter{strategy: strategy}
}

func (r *RateLimiter) Allow() bool {
	return r.strategy.Allow()
}

type RateLimiterFactory struct{}

func (f RateLimiterFactory) CreateTokenBuckelLimiter(rate int, capacity int) *RateLimiter {
	return NewRateLimiter(NewTokenBucket(rate, capacity))
}

func (f RateLimiterFactory) CreateLeakyBucketLimiter(rate int, capacity int) *RateLimiter {
	return NewRateLimiter(NewLeakyBucket(rate, capacity))
}

func (f RateLimiterFactory) CreateFixedWindowCounter(windowSize int, maxRequests int) *RateLimiter {
	return NewRateLimiter(NewFixedWindowCounter(windowSize, maxRequests))
}

func (f RateLimiterFactory) CreateSlidingWindowLog(windowSize int, maxRequests int) *RateLimiter {
	return NewRateLimiter(NewSlidingWindowLog(windowSize, maxRequests))
}

func main() {
	factory := RateLimiterFactory{}
	rateLimiter := factory.CreateTokenBuckelLimiter(3, 3)

	for i := 0; i < 10; i++ {
		if rateLimiter.Allow() {
			println("allow")
		} else {
			println("reject")
		}

		time.Sleep(500 * time.Millisecond)
	}
}
