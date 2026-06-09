package main

import (
	"testing"
	"time"
)

type fakeClock struct {
	t time.Time
}

func (fc *fakeClock) Now() time.Time            { return fc.t }
func (fc *fakeClock) Add(d time.Duration)   { fc.t = fc.t.Add(d) }

func TestRateLimiter_Allow(t *testing.T) {
	clock := &fakeClock{t: time.Unix(0, 0)}
	rl := NewRateLimiter(3, time.Second, clock)

	// Fill the window.
	for range 3 {
		if !rl.Allow() {
			t.Fatal("first 3 requests should be allowed")
		}
	}

	// Window full.
	if rl.Allow() {
		t.Fatal("4th request should be denied")
	}

	// Add past the window so old timestamps expire.
	clock.Add(2 * time.Second)
	if !rl.Allow() {
		t.Fatal("request after window should be allowed")
	}
}

func TestRateLimiter_NilClockDefaults(t *testing.T) {
	rl := NewRateLimiter(5, time.Second, nil)
	if !rl.Allow() {
		t.Fatal("first request with real clock should be allowed")
	}
}
