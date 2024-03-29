package internal

import "fmt"

// When you define a variable or function parameter as type interface{}, you are saying that it can accept any type of value.

func printValue1(v interface{}) {
	fmt.Println(v)
}

func printValue2(v any) {
	fmt.Println(v)
}
