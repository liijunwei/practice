package main

import (
	"fmt"
)

func main() {
	var a, b int

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		panic(err)
	}

	fmt.Println(a + b)
}
