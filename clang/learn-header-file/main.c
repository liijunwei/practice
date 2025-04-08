#include "max.h" // 表示 让编译器先去当前文件所在的目录找这个文件, 如果没有, 再去编译器指定的目录去找
#include <stdio.h> // 表示 让编译器去指定目录去找对应的文件

// 编译器知道自己该去哪里去找标准库的头文件
// 环境变量 和 编译器命令行参数可以指定寻找头文件的目录

int main(int argc, char const *argv[]) {
  int a = 5;
  int b = 6;

  printf("max is %f\n", max(a, b));

  return 0;
}
