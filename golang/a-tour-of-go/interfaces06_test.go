package main

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

// implements `type Stringer interface {String() string}`
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func TestI6(t *testing.T) {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	fmt.Println(a, z)

	fmt.Println(a)
	fmt.Println(z)
}
