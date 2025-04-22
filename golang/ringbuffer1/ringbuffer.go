package ringbuffer1

import "sync"

// https://medium.com/checker-engineering/a-practical-guide-to-implementing-a-generic-ring-buffer-in-go-866d27ec1a05

type RingBuffer[T any] struct {
	mu     sync.Mutex
	buffer []T
	size   int
	count  int
	write  int
}

func New[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buffer: make([]T, size),
		size:   size,
		count:  0,
		write:  0,
	}
}

// add a new item to the ring buffer
// if the buffer is full, it will overwrite the oldest item
func (b *RingBuffer[T]) Add(value T) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.buffer[b.write] = value
	b.write = (b.write + 1) % b.size

	if b.count < b.size {
		b.count++
	}
}

// return the items of the ring buffer in FIFO order
// if the buffer is empty, it will return an empty slice
// if the buffer is not full, it will return the items in the order they were added
func (b *RingBuffer[T]) Get() []T {
	b.mu.Lock()
	defer b.mu.Unlock()

	items := make([]T, 0, b.count)

	for i := range b.count {
		index := (b.write + b.size - b.count + i) % b.size
		items = append(items, b.buffer[index])
	}

	return items
}

func (b *RingBuffer[T]) Count() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.count
}
