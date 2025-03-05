package ttlcache1_test

import (
	"fmt"
	"golang-practices/ttlcache1"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSet(t *testing.T) {
	c := ttlcache1.New(ttlcache1.WithTTL(time.Second), ttlcache1.WithInterval(time.Second))

	_, ok := c.Get("a")
	require.False(t, ok)

	c.Set("a", "foo")

	v, ok := c.Get("a")
	require.True(t, ok)
	assert.Equal(t, "foo", v)
}

func TestExpiration(t *testing.T) {
	c := ttlcache1.New(ttlcache1.WithTTL(600*time.Millisecond), ttlcache1.WithInterval(500*time.Millisecond))

	_, ok := c.Get("a")
	require.False(t, ok)

	c.Set("a", "foo")

	v, ok := c.Get("a")
	require.True(t, ok)
	assert.Equal(t, "foo", v)
	assert.Equal(t, 1, c.Count())

	<-time.After(time.Second)

	_, ok = c.Get("a")
	require.False(t, ok)
	assert.Equal(t, 0, c.Count())
}

func TestTTLRenewal(t *testing.T) {
	c := ttlcache1.New(ttlcache1.WithTTL(50*time.Millisecond), ttlcache1.WithInterval(10*time.Millisecond))

	c.Set("a", "foo")

	v, ok := c.Get("a")
	require.True(t, ok)
	assert.Equal(t, "foo", v)
	assert.Equal(t, 1, c.Count())

	<-time.After(45 * time.Millisecond)
	_, ok = c.Get("a")
	require.True(t, ok)

	<-time.After(10 * time.Millisecond)
	_, ok = c.Get("a")
	require.True(t, ok)
}

func TestConcurrentAccess(t *testing.T) {
	c := ttlcache1.New(ttlcache1.WithTTL(3*time.Millisecond), ttlcache1.WithInterval(1*time.Millisecond))

	var wg sync.WaitGroup
	numWorkers := 100000
	keys := []string{"k1", "k2", "k3"}

	wg.Add(numWorkers)

	var hitcount atomic.Int64
	var misscount atomic.Int64

	for id := range numWorkers {
		go func() {
			defer wg.Done()

			key := keys[id%len(keys)]

			if rand.IntN(100) > 95 {
				c.Set(key, fmt.Sprintf("value-%d", id))
			}

			if _, ok := c.Get(key); ok {
				hitcount.Add(1)
				return
			}

			misscount.Add(1)
		}()
	}

	wg.Wait()

	// TODO add metrics to the cache struct
	t.Log("hitcount", hitcount.Load(), "misscount", misscount.Load(), "cache hit rate", float64(hitcount.Load())/float64((hitcount.Load()+misscount.Load())))

	for _, k := range keys {
		_, ok := c.Get(k)
		require.True(t, ok)
	}
}
