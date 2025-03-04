package lrucache1

import (
	"container/list"
)

// TODO support thread safety
// TODO generic key/value?

// ttl + lru cache, remove the oldest item when at capacity
type Cache struct {
	max   int                      // maximum number of entries in the LRU cache
	ll    *list.List               // use doubly linked list to maintain the "least-recently-used items"
	items map[string]*list.Element // the cache store
}

// item of the doubly linked list
type item struct {
	key   string
	value string
}

func newItem(key, value string) *item {
	return &item{
		key:   key,
		value: value,
	}
}

type CacheOption func(c *Cache)

func New(options ...CacheOption) *Cache {
	c := &Cache{
		max:   100, // default max
		ll:    list.New(),
		items: make(map[string]*list.Element),
	}

	for _, o := range options {
		o(c)
	}

	assert(c.max > 0 && c.max <= 1000)

	return c
}

func WithMaxEntries(max int) CacheOption {
	return func(c *Cache) {
		c.max = max
	}
}

func (c *Cache) Set(key, value string) {
	if lm, ok := c.items[key]; ok {
		c.ll.MoveToFront(lm)
		lm.Value.(*item).value = value

		return
	}

	newItem := c.ll.PushFront(newItem(key, value))
	c.items[key] = newItem

	if c.ll.Len() > c.max {
		c.removeOldest()
	}
}

func (c *Cache) Get(key string) (data string, ok bool) {
	if ele, ok := c.items[key]; ok {
		c.ll.MoveToFront(ele)

		itm, ok := ele.Value.(*item)
		assert(ok)

		return itm.value, true
	}

	return "", false
}

func (c *Cache) Count() int {
	return c.ll.Len()
}

func (c *Cache) removeOldest() {
	if ele := c.ll.Back(); ele != nil {
		c.removeElement(ele)
	}
}

func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)

	itm, ok := e.Value.(*item)
	assert(ok)

	delete(c.items, itm.key)
	// log.Println("evicted item:", itm.key, itm, itm.value)
}

func assert(ok bool) {
	if !ok {
		panic("assertion fails")
	}
}
