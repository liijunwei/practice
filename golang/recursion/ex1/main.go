package main

import (
	"fmt"
)

func main() {
	print("hello world")
}

func print(input string) {
	fmt.Printf("%s -> %s\n", input, reverse(input))
}

func reverse(input string) string {
	if input == "" {
		return ""
	}

	// runtime.Breakpoint()
	// fmt.Printf("input[0:1]: %v\n", input[0:1])
	return reverse(input[1:]) + input[0:1]
}
