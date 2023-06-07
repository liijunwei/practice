package main

import (
	"fmt"
	"time"
)

var colors = []string{"red", "green", "blue"}

func main() {
	fmt.Println("serial running")
	for _, c := range colors {
		fmt.Println(c)
	}

	fmt.Println("parallel running")
	for _, c := range colors {
		// c := c // fix by adding this line
		go func() {
			fmt.Println(c)
		}()
	}

	time.Sleep(3 * time.Second)
}
