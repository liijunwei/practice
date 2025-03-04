package lrucache1_test

import (
	"golang-practices/lrucache1"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSet(t *testing.T) {
	c := lrucache1.New(lrucache1.WithMaxEntries(3))
	c.Set("a", "foo")

	assert.Equal(t, 1, c.Count())

	val, ok := c.Get("a")
	require.True(t, ok)
	assert.Equal(t, "foo", val)
}

func TestEviction(t *testing.T) {
	c := lrucache1.New(lrucache1.WithMaxEntries(3))
	assert.Equal(t, 0, c.Count())

	c.Set("a", "foo1")
	c.Set("b", "foo2")
	c.Set("c", "foo3")
	c.Set("d", "foo4")

	assert.Equal(t, 3, c.Count())

	_, ok := c.Get("a")
	require.False(t, ok)

	val, ok := c.Get("b")
	require.True(t, ok)
	assert.Equal(t, "foo2", val)

	val, ok = c.Get("c")
	require.True(t, ok)
	assert.Equal(t, "foo3", val)

	val, ok = c.Get("d")
	require.True(t, ok)
	assert.Equal(t, "foo4", val)

	_, ok = c.Get("e")
	require.False(t, ok)

	c.Set("f", "foo6")
	assert.Equal(t, 3, c.Count())

	c.Set("c", "foo33")
	assert.Equal(t, 3, c.Count())
	val, ok = c.Get("c")
	require.True(t, ok)
	assert.Equal(t, "foo33", val)
}
