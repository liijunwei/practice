package main

import (
	"fmt"
	"testing"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func TestE2(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x := <-c // receive from c
	y := <-c

	fmt.Printf("partial_x: %d \npartial_y: %d \nsum: %d \n", x, y, x+y)
}
