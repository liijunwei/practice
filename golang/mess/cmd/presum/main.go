package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	print([]int{1, 1, 1, 1, 1})
	print([]int{1, 2, 3, 4, 5})
	print(parse(os.Args[1]))
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

func parse(input string) []int {
	strs := strings.Split(input, ",")
	nums := make([]int, 0, len(strs))

	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}

	return nums
}
