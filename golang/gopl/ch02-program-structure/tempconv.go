package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func init() {
	fmt.Println(1)
}
func init() {
	fmt.Println(2)
}
func init() {
	fmt.Println(3)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	fmt.Println(AbsoluteZeroC)
	fmt.Println(FreezingC)
	fmt.Println(BoilingC)

	fmt.Println(CToF(AbsoluteZeroC))
	fmt.Println(CToF(FreezingC))
	fmt.Println(CToF(BoilingC))

	f := Fahrenheit(10)
	fmt.Println(FToC(f))

	fmt.Println(AbsoluteZeroC.String())
}
