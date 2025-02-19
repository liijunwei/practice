package main

// The interface type that specifies zero methods is known as the empty interface: `interface{}`
// An empty interface may hold values of any type

import (
	"fmt"
	"testing"
)

func TestI4(t *testing.T) {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

// describe 空的接口
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
