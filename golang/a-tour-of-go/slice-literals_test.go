package main

import (
	"fmt"
	"testing"
)

func TestSliceLiteral(t *testing.T) {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	fmt.Println(q[1:3])

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
}
