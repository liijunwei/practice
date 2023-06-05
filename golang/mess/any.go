package main

import "fmt"

// When you define a variable or function parameter as type interface{}, you are saying that it can accept any type of value.

func printValue1(v interface{}) {
    fmt.Println(v)
}

func printValue2(v any) {
    fmt.Println(v)
}

func main() {
    printValue1(42)
    printValue1("Hello, world!")

    printValue1(42)
    printValue2("Hello, world!")
}
