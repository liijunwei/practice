package internal

import (
	"fmt"
	"testing"
)

const (
	ConstMonday    = 0
	ConstTuesday   = 1
	ConstWednesday = 2
)

// gofmt -w *.go
func Test_Const(t *testing.T) {
	fmt.Println("monday    ", ConstMonday)
	fmt.Println("tuesday   ", ConstTuesday)
	fmt.Println("wednesday ", ConstWednesday)
}

const (
	I1Monday = 1 << iota
	I1Tuesday
	I1Wednesday
)

func Test_Iota1(t *testing.T) {
	fmt.Println("monday    ", I1Monday)
	fmt.Println("tuesday   ", I1Tuesday)
	fmt.Println("wednesday ", I1Wednesday)
}

const (
	_        = iota
	I2Monday = 1 << iota
	I2Tuesday
	_
	I2Wednesday
)

func Test_Iota2(t *testing.T) {
	fmt.Println("monday    ", I2Monday)
	fmt.Println("tuesday   ", I2Tuesday)
	fmt.Println("wednesday ", I2Wednesday)
}
