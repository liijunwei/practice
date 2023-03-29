package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Counter struct {
	mutex sync.Mutex
	count uint64
	n     uint64
}

func (c *Counter) can_produce() bool {
	return c.count < c.n
}

func (c *Counter) can_consume() bool {
	return c.count > 0
}

func t_producer(c *Counter) {
	for {
		c.mutex.Lock()

		if c.can_produce() {
			c.count += 1
			fmt.Print("(")
		}

		c.mutex.Unlock()
	}
}

func t_consumer(c *Counter) {
	for {
		c.mutex.Lock()

		if c.can_consume() {
			c.count -= 1
			fmt.Print(")")
		}

		c.mutex.Unlock()
	}
}

func main() {
	var c Counter
	var n, _ = strconv.Atoi(os.Args[1])
	var T, _ = strconv.Atoi(os.Args[2])
	c.n = uint64(n)

	for i := 0; i < T; i++ {
		go func() {
			t_producer(&c)
		}()

		go func() {
			t_consumer(&c)
		}()
	}

	time.Sleep(3 * time.Second)
}
