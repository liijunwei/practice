package main

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestE1(t *testing.T) {

	var a [2]string

	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// fix dependency of go-spew
	spew.Dump(primes)
}
