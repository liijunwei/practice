package main

import (
	"fmt"
	"strings"
)

// go run p2/main.go < p2/input.txt
func main() {
	var count int

	n, _ := fmt.Scan(&count)
	assert(n == 1)

	// for range n {
	for i := 0; i < count; i++ {
		var a, b int

		if n, _ := fmt.Scan(&a, &b); n == 0 {
			break
		}

		fmt.Println(a + b)
	}
}

func assert(ok bool, msg ...string) {
	if !ok {
		panic("boom: " + strings.Join(msg, ","))
	}
}
