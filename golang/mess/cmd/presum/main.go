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
	prefix := make([]int, len(input)+1)

	for i := range len(input) {
		prefix[i+1] = prefix[i] + input[i]
	}

	return prefix[1:]
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
