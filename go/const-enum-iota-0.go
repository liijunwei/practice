package main

import "fmt"

const (
	monday    = 0
	tuesday   = 1
	wednesday = 2
)

// gofmt -w *.go
func main() {
	fmt.Println("monday    ", monday)
	fmt.Println("tuesday   ", tuesday)
	fmt.Println("wednesday ", wednesday)
}
