package main

// gofmt -w demo.go && go run demo.go
import (
	"fmt"
	"strings"
)

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = []string{}

	for _, val := range arr {
		newArray = append(newArray, fn(val))
	}

	return newArray
}

func Reduce(arr []string, fn func(s string) int) int {
	sum := 0

	for _, val := range arr {
		sum += fn(val)
	}

	return sum
}

func Filter(arr []int, fn func(s int) bool) []int {
	var newArray = []int{}

	for _, val := range arr {
		if fn(val) {
			newArray = append(newArray, val)
		}
	}

	return newArray
}

func main() {
	var list = []string{"nihao", "mingtian", "nihao", "shenghuo"}

	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})

	fmt.Println("MapStrToStr:", x)

	y := Reduce(list, func(s string) int {
		return len(s)
	})

	fmt.Println("Reduce:", y) // 5+8+5+8 = 26

	intset := []int{1, 2, 3, 4, 5, 6}
	odds := Filter(intset, func(n int) bool {
		return n%2 == 1
	})

	evens := Filter(intset, func(n int) bool {
		return n%2 == 0
	})

	greaterThan3 := Filter(intset, func(n int) bool {
		return n > 3
	})

	fmt.Println("Filter odds:        ", odds)
	fmt.Println("Filter evens:       ", evens)
	fmt.Println("Filter greaterThan3:", greaterThan3)
}
