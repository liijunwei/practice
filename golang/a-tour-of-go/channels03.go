package main

// https://go.dev/tour/concurrency/4
// Only the sender should close a channel, never the receiver.
// Sending on a closed channel will cause a panic. (unclear)

import (
	"fmt"
	"testing"
)

func fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TestE4(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci1(cap(c), c)

	// The loop for i := range c receives values from the channel repeatedly until it is closed.
	for i := range c {
		fmt.Println(i)
	}
}
