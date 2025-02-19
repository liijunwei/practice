package main

import (
	"fmt"
	"testing"
)

var a = 11

func foo(n int) {
	a := 1
	a += n
}

func TestShaodowing(t *testing.T) {
	fmt.Println("a = ", a)
	foo(5)
	fmt.Println("after calling foo, a = ", a)
}
