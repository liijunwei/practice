package main

// from jyy course
import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)

	const n = 45
	fibN := fibonacci(n) // slow calculation
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
