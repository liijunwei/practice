package syncmap

import "sync"

// type safe wrapper of sync.Map
type SyncMap[K comparable, V any] struct {
	sync.Map
}

func (m *SyncMap[K, V]) Load(key K) (V, bool) {
	var zeroVal V

	v, ok := m.Map.Load(key)
	if !ok {
		return zeroVal, false
	}

	return v.(V), true
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.Map.Store(key, value)
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.Map.Delete(key)
}

func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.Map.Range(func(k, v any) bool {
		key := k.(K)
		value := v.(V)

		return f(key, value)
	})
}

func (m *SyncMap[K, V]) Clear() {
	m.Range(func(k K, _ V) bool {
		m.Delete(k)

		return true
	})
}

func (m *SyncMap[K, V]) Len() int32 {
	var count int32

	m.Range(func(_ K, _ V) bool {
		count++
		return true
	})

	return count
}
