package main

import "fmt"

const (
	_      = iota
	monday = 1 << iota
	tuesday
	_
	wednesday
)

func main() {
	fmt.Println("monday    ", monday)
	fmt.Println("tuesday   ", tuesday)
	fmt.Println("wednesday ", wednesday)
}
