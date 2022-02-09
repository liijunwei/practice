/*
page 151

为了更好的掌握有关概念, 下面使用read和write构造类似于getchar/putchar等的高级函数

如果要在包含头文件 <stdio.h> 的情况下编译这些版本的getchar函数, 就有必要用 #undef预处理指令取消getchar的宏定义, 因为在头文件里, getchar是以宏方式实现的
*/

#include <stdio.h>
#include <unistd.h>

/* TOOD 问题: 为什么不用undef也没问题??? */
#undef getchar

int getchar() {
  static char buf[BUFSIZ];
  static char *bufp = buf;
  static int n = 0;

  if (n == 0) {
    n = read(0, buf, sizeof(buf));
    bufp = buf;
  }

  return (--n >= 0) ? (unsigned char) *bufp++ : EOF;
}

// gcc ch8-The-UNIX-System-Interface/read-demo04.c && echo "afternoon" | ./a.out
int main(int argc, char const *argv[]) {
  char c;

  while ((c = getchar()) != EOF) {
    putchar(c);
  }

  return 0;
}
