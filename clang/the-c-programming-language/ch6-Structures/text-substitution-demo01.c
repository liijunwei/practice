/*
page 125

编写一个表查找程序包的核心部分代码

这段代码很典型, 可以在宏处理器或者编译器的符号表管理例程中找到
例如: 遇到 `#define IN 1` 时, 需要把 `IN` 替换为 `1`

以下两个函数用来处理名字和替换文本
install(s, t)函数把名字s和替换文本t记录到某个表中, 其中s和t仅仅是字符串
lookup(s)函数在表中查找s, 若找到, 则返回指向该处的指针, 否则返回NULL

该算法采用散列查找方法

*/

#include <stdio.h>

int main(int argc, char const *argv[])
{

  return 0;
}
