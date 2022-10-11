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

func Filter(arr []int, fn func(s int) int) int {
	sum := 0

	for _, val := range arr {
		sum += fn(val)
	}

	return sum
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
}
