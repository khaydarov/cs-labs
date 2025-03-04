package main

import (
	"sync"
	"time"
)

type FixedWindowCounter struct {
	windowSize    int
	maxRequests   int
	requestsCount int
	currentWindow time.Time

	mu sync.Mutex
}

func NewFixedWindowCounter(windowSize int, maxRequests int) *FixedWindowCounter {
	return &FixedWindowCounter{
		windowSize:    windowSize,
		maxRequests:   maxRequests,
		requestsCount: 0,
		currentWindow: time.Now(),
	}
}

func (fwc *FixedWindowCounter) Allow() bool {
	fwc.mu.Lock()
	defer fwc.mu.Unlock()

	now := time.Now().Unix()
	window := now / int64(fwc.windowSize)

	if window != fwc.currentWindow.Unix()/int64(fwc.windowSize) {
		fwc.currentWindow = time.Now()
		fwc.requestsCount = 0
	}

	if fwc.requestsCount < fwc.maxRequests {
		fwc.requestsCount++
		return true
	}

	return false
}
