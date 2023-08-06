package internal

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	counter atomic.Int32
	wg      sync.WaitGroup = sync.WaitGroup{}
)

// what happened?
func TestReturnFromGoRouting(t *testing.T) {
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
