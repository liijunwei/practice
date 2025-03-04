package ttlcache1_test

import (
	"golang-practices/ttlcache1"
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
