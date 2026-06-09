package main

import (
	"testing"
	"time"
)

type fakeClock struct {
	t time.Time
}

func (fc *fakeClock) Now() time.Time { return fc.t }

func TestRateLimiter_Allow(t *testing.T) {
	clock := &fakeClock{t: time.Unix(0, 0)}
	rl := NewRateLimiter(3, time.Second, clock)

	// First 3 requests should be allowed
	for i := range 3 {
		if !rl.Allow() {
			t.Errorf("request %d should be allowed", i)
		}
	}

	// 4th request should be denied
	if rl.Allow() {
		t.Error("4th request should be denied")
	}

	// Advance clock past the window so old timestamps expire.
	clock.t = time.Unix(2, 0)
	if !rl.Allow() {
		t.Error("request in new window should be allowed")
	}
}

func TestRateLimiter_NilClockDefaults(t *testing.T) {
	rl := NewRateLimiter(5, time.Second, nil)
	// Should not panic, uses real clock
	if !rl.Allow() {
		t.Error("first request with real clock should be allowed")
	}
}
