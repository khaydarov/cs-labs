package main

import (
	"sync"
	"time"
)

type dequeu []int

func (d *dequeu) pushFront(val int) {
	*d = append([]int{val}, *d...)
}

func (d *dequeu) pushBack(val int) {
	*d = append(*d, val)
}

func (d *dequeu) popFront() int {
	val := (*d)[0]
	*d = (*d)[1:]
	return val
}

func (d *dequeu) popBack() int {
	val := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return val
}

func (d *dequeu) front() int {
	return (*d)[0]
}

func (d *dequeu) Size() int {
	return len(*d)
}

type SlidingWindowLog struct {
	windowSize  int
	maxRequests int
	requestsLog dequeu

	mu sync.Mutex
}

func NewSlidingWindowLog(windowSize int, maxRequests int) *SlidingWindowLog {
	return &SlidingWindowLog{
		windowSize:  windowSize,
		maxRequests: maxRequests,
		requestsLog: dequeu{},
	}
}

func (swl *SlidingWindowLog) Allow() bool {
	swl.mu.Lock()
	defer swl.mu.Unlock()

	now := time.Now().Unix()

	for swl.requestsLog.Size() > 0 && int(now)-swl.requestsLog.front() >= swl.windowSize {
		swl.requestsLog.popFront()
	}

	if len(swl.requestsLog) < swl.maxRequests {
		swl.requestsLog.pushBack(int(now))
		return true
	}

	return false
}
