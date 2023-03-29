package main

import (
	"fmt"
	"time"
)

var counter = 0

// WIP
func t_producer() {
	for {
		fmt.Print("(")
		counter += 1
	}
}

func t_consumer() {
	for {
		fmt.Print(")")
		counter -= 1
	}
}

func main() {
	go func() {
		t_producer()
	}()

	go func() {
		t_producer()
	}()

	go func() {
		t_consumer()
	}()

	time.Sleep(10 * time.Second)
}
