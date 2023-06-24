package main

import "fmt"

func f() {}

var g = "g"

func main() {
	f := "f" // shadows package level func f
	fmt.Println(f)
	fmt.Println(g)
}
