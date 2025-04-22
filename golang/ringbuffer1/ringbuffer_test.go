package ringbuffer1

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRingBuffer(t *testing.T) {
	t.Parallel()

	t.Run("single thread", func(t *testing.T) {
		t.Parallel()

		b := New[int](5)
		b.Add(1)
		b.Add(2)
		b.Add(3)

		assert.Equal(t, []int{1, 2, 3}, b.Get())

		b.Add(4)
		b.Add(5)
		b.Add(6)

		assert.Equal(t, []int{2, 3, 4, 5, 6}, b.Get())

		b.Add(7)
		b.Add(8)

		assert.Equal(t, []int{4, 5, 6, 7, 8}, b.Get())
	})

	t.Run("parallel", func(t *testing.T) {
		t.Parallel()

		b := New[int](3)
		var wg sync.WaitGroup

		addValues := func(values []int) {
			for _, value := range values {
				b.Add(value)
				<-time.After(10 * time.Millisecond) // Simulate delay
			}
			wg.Done()
		}

		readValues := func() {
			prices := b.Get()
			require.True(t, len(prices) >= 0 && len(prices) <= b.size)
			wg.Done()
		}

		wg.Add(3)
		go addValues([]int{1, 2, 3})
		go addValues([]int{4, 5})
		go addValues([]int{6, 7, 8})

		wg.Add(2)
		go readValues()
		go readValues()

		wg.Wait()

		finalValues := b.Get()

		for _, value := range finalValues {
			require.True(t, value >= 1 && value <= 8)
		}

		assert.Equal(t, b.size, len(finalValues))
	})
}
