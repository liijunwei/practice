package main

import (
	"fmt"
	"testing"
)

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
// func (t T) M() {
// 	fmt.Println(t.S)
// }

// This method means type *T implements the interface I,
func (t *T) M() {
	fmt.Println(t.S)
}

func TestE7(t *testing.T) {
	var i I = &T{"hello"}
	i.M()
}
