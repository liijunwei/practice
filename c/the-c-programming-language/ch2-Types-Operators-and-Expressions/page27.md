+ C语言提供了以下几种*基本数据类型*
    + char   字符型 占一个字节(8bit), 可以存放字符集中的一个字符
    + int    整型, 通常反映了所用机器中整数的最自然长度
    + float  单精度浮点型
    + double 双精度浮点型

+ 此外可以在基本数据类型前面加限定符
    + short 通常为16位
    + long  通常为32位
    + int   可以为 16位或者32位
```c
short int a;
long int b

// 这种情况下, int 可以省略
```

+ 类型长度的限制
    + short/int至少为16位
    + long至少为32位
    + short不得长于int, int不得长于long

+ 类型限定符 signed/unsigned 可以用于限定char类型或者任意整型

+ 相关的类型长度定义参考: <limits.h> <float.h>
