package ttlcache2_test

import (
	"encoding/json"
	"fmt"
	"golang-practices/ttlcache2"
	"math/rand/v2"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSet(t *testing.T) {
	t.Parallel()

	t.Run("ttl", func(t *testing.T) {
		t.Parallel()

		ctx := t.Context()

		// default ttl is 3s
		cache := ttlcache2.New[string, int]()
		cache.Set(ctx, "key", 1)

		item, ok := cache.Get(ctx, "key")
		require.True(t, ok)
		assert.Equal(t, 1, item)

		<-time.After(ttlcache2.DefaultTTL + time.Millisecond*10)
		_, ok = cache.Get(ctx, "key")
		require.False(t, ok)
	})

	t.Run("Delete", func(t *testing.T) {
		t.Parallel()

		ctx := t.Context()

		// default ttl is 3s
		cache := ttlcache2.New[string, int]()
		cache.Set(ctx, "key", 1)

		item, ok := cache.Get(ctx, "key")
		require.True(t, ok)
		assert.Equal(t, 1, item)

		cache.Delete(ctx, "key")
		_, ok = cache.Get(ctx, "key")
		require.False(t, ok)
	})

	t.Run("Fetch", func(t *testing.T) {
		t.Parallel()

		ctx := t.Context()

		cache := ttlcache2.New[string, int]()
		_, ok := cache.Get(ctx, "key")
		require.False(t, ok)

		item, err := cache.Fetch(ctx, "key", func() (int, error) {
			return 50501, nil // function to fetch the value
		})
		require.NoError(t, err)
		assert.Equal(t, 50501, item)

		v, ok := cache.Get(ctx, "key")
		require.True(t, ok)
		assert.Equal(t, 50501, v)
	})
}

func BenchmarkFetch(b *testing.B) {
	ctx := b.Context()

	cache := ttlcache2.New[string, int32]()
	b.ResetTimer()

	for b.Loop() {
		num := rand.Int32()
		_, err := cache.Fetch(ctx, "key", func() (int32, error) {
			return num, nil
		})
		require.NoError(b, err)
	}
}

func TestConcurrentAccess(t *testing.T) {
	cache := ttlcache2.New[string, string]()

	ctx := t.Context()

	var wg sync.WaitGroup
	numWorkers := 1000000
	keys := []string{"k1", "k2", "k3", "k4", "k5"}

	wg.Add(numWorkers)

	for id := range numWorkers {
		go func() {
			defer wg.Done()

			key := keys[id%len(keys)]

			_, err := cache.Fetch(ctx, key, func() (string, error) {
				return fmt.Sprintf("value-%d", id), nil
			})
			if err != nil {
				t.Log("unexpected error:", err)
			}
		}()
	}

	wg.Wait()

	data, err := json.Marshal(cache.MetricStat())
	require.NoError(t, err)
	t.Log("cache stat", string(data))
}
