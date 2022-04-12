package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func demo(x, y, z string) (string, string, string) {
	return y, z, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	ljw, wxy, lwj := demo("ljw", "wxy", "lwj")
	fmt.Println(ljw, wxy, lwj)
}
