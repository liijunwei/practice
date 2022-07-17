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


