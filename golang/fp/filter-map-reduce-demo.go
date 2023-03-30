package main

// chatGPT generated
import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// Map example: multiply each number by 2
	mappedNums := Map(nums, func(n int) int {
		return n * 2
	})
	fmt.Println(mappedNums) // [2 4 6 8 10]

	// Filter example: keep only even numbers
	filteredNums := Filter(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(filteredNums) // [2 4]

	// Reduce example: sum all the numbers
	sum := Reduce(nums, func(acc, n int) int {
		return acc + n
	})
	fmt.Println(sum) // 15
}

// Map function: applies a function to each element of a slice and returns a new slice
func Map(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = fn(n)
	}
	return result
}

// Filter function: keeps only elements of a slice that satisfy a given condition and returns a new slice
func Filter(nums []int, fn func(int) bool) []int {
	result := make([]int, 0)
	for _, n := range nums {
		if fn(n) {
			result = append(result, n)
		}
	}
	return result
}

// Reduce function: applies a function to all elements of a slice and returns a single result
func Reduce(nums []int, fn func(int, int) int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = fn(result, nums[i])
	}
	return result
}
