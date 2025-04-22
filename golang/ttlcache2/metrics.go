package ttlcache2

import (
	"sync/atomic"
)

type CacheMetrics struct {
	name         string
	hitCounter   atomic.Int64
	missCounter  atomic.Int64
	errorCounter atomic.Int64
}

func NewCacheMetrics(name string) *CacheMetrics {
	return &CacheMetrics{
		name:         name,
		hitCounter:   atomic.Int64{},
		missCounter:  atomic.Int64{},
		errorCounter: atomic.Int64{},
	}
}

func (c *CacheMetrics) Hit() {
	c.hitCounter.Add(1)
}

func (c *CacheMetrics) Miss() {
	c.missCounter.Add(1)
}

func (c *CacheMetrics) Error() {
	c.errorCounter.Add(1)
}

func (c *CacheMetrics) Stat() map[string]any {
	hitCount := c.hitCounter.Load()
	missCount := c.missCounter.Load()
	errorCount := c.errorCounter.Load()

	stat := make(map[string]any)

	stat["name"] = c.name
	stat["hit"] = hitCount
	stat["miss"] = missCount
	stat["error"] = errorCount

	total := hitCount + missCount + errorCount
	stat["total"] = hitCount + missCount + errorCount

	stat["hit_rate"] = float64(hitCount) / float64(total)
	stat["miss_rate"] = float64(missCount) / float64(total)
	stat["error_rate"] = float64(errorCount) / float64(total)

	return stat
}
