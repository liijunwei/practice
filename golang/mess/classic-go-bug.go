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
		go func() {
			fmt.Println(c)
		}()
	}

	time.Sleep(3 * time.Second)
}
