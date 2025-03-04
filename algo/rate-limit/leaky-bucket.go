package main

import (
	"sync"
	"time"

	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
)

type LeakyBucket struct {
	rate     int
	capacity int
	bucket   *llq.Queue
	lastLeak time.Time

	mu sync.Mutex
}

func NewLeakyBucket(rate, capacity int) *LeakyBucket {
	return &LeakyBucket{
		rate:     rate,
		capacity: capacity,
		bucket:   llq.New(),
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(lb.lastLeak)
	lb.lastLeak = now

	leaked := int(elapsed.Seconds() * float64(lb.rate))
	if leaked > 0 {
		for i := 0; i < leaked; i++ {
			if lb.bucket.Size() > 0 {
				lb.bucket.Dequeue()
			}
		}
		lb.lastLeak = now
	}

	if lb.bucket.Size() < lb.capacity {
		lb.bucket.Enqueue(now)
		return true
	}

	return false
}
