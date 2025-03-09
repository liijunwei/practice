package main

import "fmt"

func main() {
	print([]int{1, 1, 1, 1, 1})
	print([]int{1, 2, 3, 4, 5})
}

func print(input []int) {
	fmt.Println("=====")
	fmt.Println(input)
	fmt.Println(presum(input))
}

func presum(input []int) []int {
	prefix := make([]int, len(input))

	sum := 0

	for i := 0; i < len(input); i++ {
		sum += input[i]
		prefix[i] = sum
	}

	return prefix
}
