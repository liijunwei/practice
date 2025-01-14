package main

import (
	"fmt"
)

// go run p3/main.go < p3/input.txt
func main() {
	for {
		var a, b int

		if n, _ := fmt.Scan(&a, &b); n == 0 || (a == 0 && b == 0) {
			break
		}

		fmt.Println(a + b)
	}
}
