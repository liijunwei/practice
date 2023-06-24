package main

import "fmt"

func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

// the new(T) creates an unnamed variable of type T, initializes it to the zero value of T, and returns its address, which is a value of type *T
func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 999
	fmt.Println(*p)
	fmt.Println(*newInt1())
	fmt.Println(*newInt2())

	i, j, k := 1, 2, 3
	fmt.Println(i, j, k)

	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(medals)
	medals[0] = "gold1"
	fmt.Println(medals)
}
