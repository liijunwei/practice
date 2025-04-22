package ttlcache2

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"golang-practices/syncmap"

	"golang.org/x/sync/singleflight"
)

var DefaultTTL = time.Second * 3

type Cache[K comparable, V any] interface {
	Fetch(ctx context.Context, key K, fetch func() (V, error)) (V, error)
	Get(ctx context.Context, key K) (V, bool)
	Set(ctx context.Context, key K, value V)
	Delete(ctx context.Context, key K)
}

type MemCache[K comparable, V any] struct {
	store syncmap.SyncMap[K, item[V]]
	ttl   time.Duration
	sg    singleflight.Group
}

type item[V any] struct {
	value   V
	addedAt time.Time
}

var _ Cache[struct{}, struct{}] = &MemCache[struct{}, struct{}]{}

// TODO add configurable options(TTL, etc.)
func NewMemCache[K comparable, V any]() *MemCache[K, V] {
	cache := &MemCache[K, V]{
		store: syncmap.SyncMap[K, item[V]]{},
		sg:    singleflight.Group{},
		ttl:   DefaultTTL,
	}

	return cache
}

func (m *MemCache[K, V]) Get(ctx context.Context, key K) (V, bool) {
	var zeroVal V

	if item, ok := m.store.Load(key); ok {
		if m.ttl > 0 && time.Since(item.addedAt) > m.ttl {
			m.store.Delete(key)

			return zeroVal, false
		}

		return item.value, true
	}

	return zeroVal, false
}

func (m *MemCache[K, V]) Set(_ context.Context, key K, value V) {
	val := item[V]{
		value:   value,
		addedAt: time.Now(),
	}
	m.store.Store(key, val)
}

func (m *MemCache[K, V]) Delete(_ context.Context, key K) {
	m.store.Delete(key)
}

func (m *MemCache[K, V]) Fetch(
	ctx context.Context,
	key K,
	fetch func() (V, error),
) (V, error) {
	if item, ok := m.Get(ctx, key); ok {
		return item, nil
	}

	var zeroVal V

	singleflightkey, err := json.Marshal(key)
	if err != nil {
		return zeroVal, fmt.Errorf("failed to marshal key: %w", err)
	}

	// TODO understand why using singleflight here
	// what happens if it's not used
	value, err, _ := m.sg.Do(string(singleflightkey), func() (interface{}, error) {
		item, err := fetch()
		if err != nil {
			return zeroVal, err
		}

		m.Set(ctx, key, item)

		return item, nil
	})

	if err != nil {
		return zeroVal, fmt.Errorf("failed to fetch value: %w", err)
	}

	item, ok := value.(V)
	if !ok {
		return zeroVal, fmt.Errorf("failed to cast value to expected type")
	}

	return item, nil
}
