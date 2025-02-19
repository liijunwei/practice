// https://go.dev/tour/methods/1

package main

// Go does not have classes. However, you can define methods on types.
// go没有类, 但可以给用struc定义的类型定义方法

import (
	"fmt"
	"math"
	"testing"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func TestM1(t *testing.T) {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v2 := Vertex{1, 1}
	fmt.Println(v2.Abs())
}
