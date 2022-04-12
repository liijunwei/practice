package main

import "fmt"

const (
	monday = 1 << iota
	tuesday
	wednesday
)

func main() {
	fmt.Println("monday    ", monday)
	fmt.Println("tuesday   ", tuesday)
	fmt.Println("wednesday ", wednesday)
}
