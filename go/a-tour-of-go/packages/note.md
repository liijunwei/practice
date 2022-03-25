所有的go程序都由包(package)组成

main包里的程序时整个程序的入口

约定如下: 包的名称是import路径的最后一级的名称

In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

A function can return any number of results.

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

```go
package main

import "fmt"

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println(split(3))
}

// 1 2
```

Outside a function, every statement begins with a keyword (var, func, and so on) and so the `:=` construct is not available.

+ Go's basic types are
    + bool
    + int int8 int16 int32 int64
    + uint uint8 uint16 uint32 uint64 uintptr
    + byte // alias for uint8
    + rune // alias for int32, represents a unicode code point(TODO what is unicode code point?)
    + float32 float64
    + complex64 complex128

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

+ Variables declared without an explicit initial value are given their zero value.
    + 0 for numeric types,
    + false for the boolean type, and
    + "" (the empty string) for strings.

Unlike in C, in Go assignment between items of different type requires an explicit conversion.

```go
i := 42
f := float64(i)
u := uint(f)
```

+ TODO https://go.dev/tour/basics/14


