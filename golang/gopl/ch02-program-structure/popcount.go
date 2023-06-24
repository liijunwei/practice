package main

import "fmt"

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	fmt.Println(PopCount(1))
	fmt.Println(PopCount(2))
	fmt.Println(PopCount(4))

	//var b byte = 100
	var a uint64 = 4
	fmt.Println(a << 1)
}
