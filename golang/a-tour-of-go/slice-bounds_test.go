package main

// For the array
//     var a [10]int

// these slice expressions are equivalent:

//     a[0:10]
//     a[:10]
//     a[0:]
//     a[:]

import (
	"fmt"
	"testing"
)

func TestSliceBoundary(t *testing.T) {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
