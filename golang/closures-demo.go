package main

import "fmt"

func intSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextIntegerFunc1 := intSequence()

	fmt.Println(nextIntegerFunc1())
	fmt.Println(nextIntegerFunc1())
	fmt.Println(nextIntegerFunc1())
	fmt.Println()

	nextIntegerFunc2 := intSequence()
	fmt.Println(nextIntegerFunc2())
}
