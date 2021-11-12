package main

import "fmt"

const (
	_ = iota
	monday = 1 << iota
	tuesday
	_
	wednesday
)

// gofmt -w const-enum-iota-2.go
func main() {
	fmt.Println("monday    ", monday)
	fmt.Println("tuesday   ", tuesday)
	fmt.Println("wednesday ", wednesday)
}
