package main

// buffered channel

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	// ch <- 3 // Modify the example to overfill the buffer and see what happens.

	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
}
