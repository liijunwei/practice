package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Counter struct {
	m     sync.Mutex
	count uint64
}

func (c *Counter) t_producer(n uint64) {
	for {
		c.m.Lock()

		if c.can_produce(n) {
			c.count += 1
			fmt.Print("(")
		}

		c.m.Unlock()
	}
}

func (c *Counter) t_consumer() {
	for {
		c.m.Lock()

		if c.can_consume() {
			c.count -= 1
			fmt.Print(")")
		}

		c.m.Unlock()
	}
}

func (c *Counter) can_produce(n uint64) bool {
	return c.count < n
}

func (c *Counter) can_consume() bool {
	return c.count > 0
}

func main() {
	var c Counter
	var n, _ = strconv.Atoi(os.Args[1])
	var T, _ = strconv.Atoi(os.Args[2])

	for i := 0; i < T; i++ {
		go func() {
			c.t_producer(uint64(n))
		}()

		go func() {
			c.t_consumer()
		}()
	}

	time.Sleep(3 * time.Second)
}
