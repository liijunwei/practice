package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter atomic.Int32
	wg      sync.WaitGroup = sync.WaitGroup{}
)

func main() {
	wg.Add(1)
	go foo()

	wg.Wait()
}

func foo() {
	defer wg.Done()

	for {
		if counter.Load() > 10 {
			return
		}

		counter.Add(1)
		fmt.Println("incr...", counter.Load())
		time.Sleep(100 * time.Millisecond)
	}
}
