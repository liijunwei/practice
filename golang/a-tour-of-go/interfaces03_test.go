package main

import (
	"fmt"
	"math"
	"testing"
)

type I3 interface {
	M()
}

type T3 struct {
	S string
}

func (t *T3) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func TestI3(t *testing.T) {
	var i I3

	i = &T3{"Hello"}
	describe(i)
	i.M()

	fmt.Printf("===============\n")

	i = F(math.Pi)
	describe(i)
	i.M()
}
