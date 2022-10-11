package main

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

func main() {
	var list = []string{"nihao", "mingtian", "nihao", "shenghuo"}

	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})

	fmt.Println(x)
}
