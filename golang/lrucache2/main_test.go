package main

import (
	"fmt"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func TestLRUCache(t *testing.T) {
	size := 10
	cache := NewLRUCache(size)

	checkFixedSize := func(key string, value int) bool {
		for i := 0; i < size*3; i++ {
			cache.Put(fmt.Sprintf("%d", i), fmt.Sprintf("%d", i))
		}

		return cache.Len() == size
	}

	err := quick.Check(checkFixedSize, &quick.Config{
		MaxCount: 10000,
	})
	require.NoError(t, err)
}
