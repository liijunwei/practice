package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		panic(err)
	}

	arr := make([]string, 0, count)

	for i := 0; i < count; i++ {
		var s string
		_, err := fmt.Scan(&s)
		boom(err)

		arr = append(arr, s)
	}

	sort.Strings(arr)
	fmt.Println(strings.Join(arr, " "))
}

func boom(err error) {
	if err != nil {
		panic(err)
	}
}
