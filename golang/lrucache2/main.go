package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	cache := NewLRUCache(2)
	cache.Put("1", "1")
	cache.Put("2", "2")
	cache.Put("3", "3")
	// spew.Dump(cache)
	cache.Get("2")
	spew.Dump(cache.store)
}

type Node struct {
	key   string
	value string
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	store    map[string]*Node
}

func NewLRUCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		store:    make(map[string]*Node),
	}
}

func (c *LRUCache) Get(key string) (string, error) {
	node, ok := c.store[key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}
	c.moveToHead(node)

	return node.value, nil
}

func (c *LRUCache) Put(key, value string) {
	node, ok := c.store[key]
	if ok {
		node.value = value
		c.moveToHead(node)
		return
	}

	if len(c.store) >= c.capacity {
		c.removeTail()
	}

	node = &Node{
		key:   key,
		value: value,
	}
	c.store[key] = node
	c.addToHead(node)
}

func (c *LRUCache) Len() int {
	return len(c.store)
}

func (c *LRUCache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) removeTail() {
	node := c.tail.prev
	c.removeNode(node)
	delete(c.store, node.key)
}

func (c *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) addToHead(node *Node) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}
