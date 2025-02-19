package main

import (
	"fmt"
	"testing"
)

type I2 interface {
	M()
}

type T2 struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
// func (t T) M() {
// 	fmt.Println(t.S)
// }

// This method means type *T implements the interface I,
func (t *T2) M() {
	fmt.Println(t.S)
}

func TestE7(t *testing.T) {
	var i I2 = &T2{"hello"}
	i.M()
}
