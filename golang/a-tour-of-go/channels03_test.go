package main

// https://go.dev/tour/concurrency/4
// Only the sender should close a channel, never the receiver.
// Sending on a closed channel will cause a panic. (unclear)

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

// it's ok to read from closed channel
// it panics to send to closed channel
// so close is ONLY done by sender, not the receiver
func Test_close_channel(t *testing.T) {
	a := make(chan int)
	close(a)
	<-a
	assert.True(t, true)

	b := make(chan int)
	close(b)
	assert.Panics(t, func() { b <- 1 })
}

// make for array/slice/channel
// accept Type, not value
// return the value of type
// new for any type, it returns pointer to the type
func Test_make_new(t *testing.T) {
	tmp := make([]any, 0, 100)
	t1 := new(map[string]string)
	t2 := new([]string)
	t3 := new(int)
	t4 := new(float32)
	_ = new(int)
	_ = new(string)

	tmp = append(tmp, t1)
	tmp = append(tmp, t2)
	tmp = append(tmp, t3)
	tmp = append(tmp, t4)
}
