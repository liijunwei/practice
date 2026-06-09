package main

import (
	"sync"
	"time"
)

// Clock is an interface for getting the current time, making RateLimiter testable.
type Clock interface {
	Now() time.Time
	Add(d time.Duration)
}

type realClock struct{}

func (realClock) Now() time.Time           { return time.Now() }
func (realClock) Add(time.Duration) {} // no-op: can't advance real time

// RateLimiter implements a sliding-window rate limiter.
type RateLimiter struct {
	mu      sync.Mutex
	clock   Clock
	limit   int
	window  time.Duration
	history []time.Time
}

// NewRateLimiter creates a RateLimiter allowing limit requests per window.
func NewRateLimiter(limit int, window time.Duration, clock Clock) *RateLimiter {
	if clock == nil {
		clock = realClock{}
	}
	return &RateLimiter{
		clock:  clock,
		limit:  limit,
		window: window,
	}
}

// Allow returns true if a request is allowed.
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := rl.clock.Now()
	cutoff := now.Add(-rl.window)

	// Evict expired timestamps.
	i := 0
	for i < len(rl.history) && rl.history[i].Before(cutoff) {
		i++
	}
	rl.history = rl.history[i:]

	if len(rl.history) < rl.limit {
		rl.history = append(rl.history, now)
		return true
	}
	return false
}
