// https://go.dev/tour/methods/4

// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

// 方法接收者用指针表示, 可以修改结构里的数据

// Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.

// 如果把Scale函数声明里的 *去掉, 实际上是操作结构的副本, 不会影响结构的值

// The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

package main

import (
	"fmt"
	"math"
	"testing"
)

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestMethodPointer(t *testing.T) {
	v := Vertex2{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
