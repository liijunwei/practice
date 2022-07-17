+ 变量声明的意义在于 告诉编译器该变量可以操作的内存的边界信息, 而这种边界信息通常由变量的类型信息提供

+ go中通用的变量声明方法
```go
var a int = 10

var a int
```

+ 如果声明时, 不给变量初始化值, 默认为类型的零值
    + int -> 0
    + float -> 0.0
    + boolean -> false
    + string -> ""
    + pointer/interface/slice/channel/map/function -> nil
    + 符合类型 -> 组成元素都为零值

+ 变量声明块
```go
var (
  a int = 128
  b int8 = 6
  s string = "hello"
  c rune = 'A'
  t bool = true
)
```

+ 同一行内声明
```go
var a, b, c int = 1, 3, 5
```

+ 省略类型信息的声明
```go
var b = 13 // 类型推断为int
```

+ 如果不接受默认的类型推断结果
```go
var b = int32(13)
```

+ 语法糖
```go
var a, b, c = 13, 'A', "hello"
```

+ 段变量声明
```go
a := 13
b := 'A'
c := "Hello"

a, b, c := 13, 'A', "hello"
```

+ go语言的两类变量
    1. 包级变量(package variable)
    2. 局部变量(local variable)

+ 包级变量
    + 包级变量只能用带有var关键字的变量声明形式, 不能使用短变量声明形式, 但在形式细节上可以有一定的灵活度(问题: 什么意思?)
    + 声明并同时显式初始化
    + 声明但延迟初始化

+ 局部变量
    + 












