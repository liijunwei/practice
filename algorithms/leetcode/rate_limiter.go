package main

import "time"

// Clock is an interface for getting the current time, making RateLimiter testable.
type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time { return time.Now() }

// RateLimiter implements a sliding-window rate limiter.
type RateLimiter struct {
	clock   Clock
	limit   int
	window  time.Duration
	history []time.Time
}

// NewRateLimiter creates a RateLimiter with the given limit per window.
// The clock defaults to the real system clock if nil is passed.
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

// Allow returns true if a request is allowed, false if the rate limit is exceeded.
func (rl *RateLimiter) Allow() bool {
	now := rl.clock.Now()
	cutoff := now.Add(-rl.window)

	// Evict all requests older than the window.
	i := 0
	for i < len(rl.history) && rl.history[i].Before(cutoff) {
		i++
	}
	rl.history = rl.history[i:]

	// If capacity has ballooned (e.g. after a burst), shrink the backing array.
	if cap(rl.history) > 4*len(rl.history) {
		trimmed := make([]time.Time, len(rl.history))
		copy(trimmed, rl.history)
		rl.history = trimmed
	}

	if len(rl.history) < rl.limit {
		rl.history = append(rl.history, now)
		return true
	}
	return false
}
