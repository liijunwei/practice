package main

import (
	"fmt"
	"testing"
	"time"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func TestE5(t *testing.T) {
	go say("world")
	say("hello")
}
