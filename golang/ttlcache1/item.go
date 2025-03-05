package ttlcache1

import (
	"sync"
	"time"
)

type Item struct {
	mu        sync.RWMutex
	data      string
	expiresAt time.Time
	ttl       time.Duration
}

func NewItem(data string, ttl time.Duration) *Item {
	return &Item{
		data:      data,
		expiresAt: time.Now().Add(ttl),
		ttl:       ttl,
	}
}

func (i *Item) renewTTL() {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.expiresAt = time.Now().Add(i.ttl)
}

func (i *Item) isExpired() bool {
	i.mu.RLock()
	defer i.mu.RUnlock()

	assert(!i.expiresAt.IsZero())

	return i.expiresAt.Before(time.Now())
}

func assert(ok bool) {
	if !ok {
		panic("assertion fails")
	}
}
