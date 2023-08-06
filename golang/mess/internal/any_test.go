package internal

import "testing"

// When you define a variable or function parameter as type interface{}, you are saying that it can accept any type of value.

func Test_any(t *testing.T) {
	printValue1(42)
	printValue1("Hello, world!")

	printValue1(42)
	printValue2("Hello, world!")
}
