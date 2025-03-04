package ttlcache1

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	ttl      time.Duration // items ttl
	interval time.Duration // cleanup interval
	items    map[string]*Item
}

type CacheOption func(c *Cache)

// TODO it's better to limit the items count
// assume cleanup interval >= 500ms
func New(options ...CacheOption) *Cache {
	c := &Cache{
		ttl:      5 * time.Second,
		interval: time.Second,
		items:    make(map[string]*Item),
	}

	for _, o := range options {
		o(c)
	}

	assert(c.ttl > 0)
	assert(c.interval >= 500*time.Millisecond)

	go c.backgroundCleanup()

	return c
}

func WithTTL(ttl time.Duration) CacheOption {
	return func(c *Cache) {
		c.ttl = ttl
	}
}

func WithInterval(interval time.Duration) CacheOption {
	return func(c *Cache) {
		c.interval = interval
	}
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item := NewItem(value, c.ttl)
	c.items[key] = item
}

func (c *Cache) Get(key string) (data string, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.items[key]
	if !ok || item.isExpired() {
		return "", false
	}

	item.touch()

	return item.data, true
}

func (c *Cache) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.items)
}

func (c *Cache) backgroundCleanup() {
	ticker := time.NewTicker(c.interval)

	for range ticker.C {
		c.cleanup()
	}
}

func (c *Cache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, item := range c.items {
		if item.isExpired() {
			delete(c.items, key)
		}
	}
}
