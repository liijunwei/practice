package internal

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func Test_AtomaticCounter(t *testing.T) {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				atomic.AddUint64(&ops, 1)
				// ops++
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
