/*
page 136

以实现函数printf的简单版本为例, 介绍如何以可移植的方式编写可处理变长参数表的函数
因为重点在于参数的处理, 所以函数minprintf只处理格式字符串和参数, 格式转换通过调用printf实现

函数printf的正确声明形式为 `int printf(char *fmt, ...)`
其中省略号表示参数表中的参数的数量和类型是可变的, 省略号只能放在参数表的末尾

因为minprintf函数不需要向printf函数那样返回实际输出的字符个数, 因此声明为 `void minprintf(char *fmt, ...)`
编写minprintf函数的关键在于如何处理甚至连名字都没有的参数表
标准头文件<stdarg.h>中定义了一组宏定义, 他们对如何遍历参数表进行了定义; 该头文件的实现因不同的机器而不同, 但提供的接口是一致的


*/

#include <stdio.h>

int main(int argc, char const *argv[]) {


  return 0;
}
