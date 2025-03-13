package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// e.g.
// go run mess/cmd/presum/main.go 1,2,3,4,9 1 4
//
// raw:    [1 2 3 4 9]
// presum: [1 3 6 10 19]
// sum(1,4) = prefix[4] - prefix[1-1] = prefix[4] - prefix[0] = 18
func main() {
	print([]int{1, 1, 1, 1, 1})
	print([]int{1, 2, 3, 4, 5})
	print(parse(os.Args[1]))

	left := mustParseInt(os.Args[2])
	right := mustParseInt(os.Args[3])
	fmt.Println(rangesum(parse(os.Args[1]), left, right))
}

// range sum: sum(L, R) = prefix[R] - prefix[L-1]
// e.g. rangesum([]int{1, 1, 1, 1, 1}, 2, 3) => 1+1 = presum[3]-presum[1])
// e.g. rangesum([]int{1, 2, 3, 4, 5}, 2, 3) => 3+4 = presum[3]-presum[1])
func rangesum(input []int, left, right int) int {
	assert(left >= 0)
	assert(right >= 0)
	assert(right >= left)
	assert(right < len(input))

	presum := presum(input)

	if left == 0 {
		return presum[right]
	}

	return presum[right] - presum[left-1]
}

func assert(ok bool) {
	if !ok {
		panic("assertion failed")
	}
}

func print(input []int) {
	fmt.Println("=====")
	fmt.Println(input)
	fmt.Println(presum(input))
}

// time complexity: O(n)
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
		nums = append(nums, mustParseInt(str))
	}

	return nums
}

func mustParseInt(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return num
}
