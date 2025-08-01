package main

import "fmt"

func main() {
	gen := naturals()
	size := 100
	lst := make([]int, 0, size)
	for range size {
		lst = append(lst, <-gen)
	}
	fmt.Println(lst)
}

// natural generator in go
func naturals() <-chan int {
	ch := make(chan int)
	go func() {
		x := 0
		for {
			ch <- x
			x++
		}
	}()
	return ch
}
