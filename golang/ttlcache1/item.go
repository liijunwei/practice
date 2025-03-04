package ttlcache1

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Item struct {
	sync.RWMutex
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

func (i *Item) touch() {
	i.Lock()
	defer i.Unlock()

	i.expiresAt = time.Now().Add(i.ttl)
}

func (i *Item) isExpired() bool {
	i.RLock()
	defer i.RUnlock()

	assert(!i.expiresAt.IsZero())

	return i.expiresAt.Before(time.Now())
}

func assert(ok bool) {
	if !ok {
		_, filename, line, _ := runtime.Caller(1)
		msg := fmt.Sprintf("assertion fails: %s:%d", filename, line)
		panic(msg)
	}
}
