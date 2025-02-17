package main

import (
	"fmt"
	"time"
)

// go run -gcflags -m cmd/tmp1/main.go
func main() {
	done := make(chan bool)

	time.NewTicker(1 * time.Second)

	// close(done)
	go func() {
		done <- true
	}()

	fmt.Println(<-done)

	// done <- true
	f1()
}

func f1() map[string]string {
	m := make(map[string]string)
	m["a"] = "a"
	return m
}
