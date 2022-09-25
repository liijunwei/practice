#include <stdio.h>

/*
page 99

命令行参数
---------------
在支持C语言的环境中, 可在程序开始执行时, 将命令行参数传递给程序
调用主函数main时, 它带有两个参数:
第一个参数(argc)的值表示运行程序时命令行中参数的数目;
第二个参数(argv)是一个指向字符串数组 的指针, 其中每个字符串对应一个参数

我们通常用多级指针处理这些字符串(TODO ???)

最简单的例子是程序 echo
echo hello, world

按照C语言的约定, argv[0]的值是启动该程序的程序名, 因此argc的值至少为1
如果argc的只为1, 则说明程序名后面没有命令行参数

ANSI标准要求argv[argc]的值必须为一个空指针
*/

// gcc -g ch5-Pointers-and-Arrays/echo-demo.c -o echo-demo && ./echo-demo hello,
// world run ch5-Pointers-and-Arrays/echo-demo.c hello world
int main(int argc, char const *argv[]) {
  printf("argc: %d\n", argc);
  printf("argv[0]: %s\n", argv[0]);
  printf("argv[1]: %s\n", argv[1]);
  printf("argv[2]: %s\n", argv[2]);
  printf("argv[3]: %s\n", argv[3]);
  printf("argv[4]: %s\n", argv[4]);
  printf("argv[5]: %s\n", argv[5]);
  printf("argv[argc]: %s\n", argv[argc]);
  printf("==============================================\n\n");

  int i;
  for (i = 1; i < argc; i++) {
    char *seperator = (i < argc - 1) ? " " : "";
    printf("%s%s", argv[i], seperator);
  }

  printf("\n");

  return 0;
}
