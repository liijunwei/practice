/*
page 136

编写一个程序, 以合理的方式打印任何输入; 该程序至少能够根据用户的习惯以8进制或16进制打印非图形字符, 并截断长文本行

https://www.runoob.com/cprogramming/c-function-iscntrl.html
C 库函数 void iscntrl(int c) 检查所传的字符是否是控制字符。
根据标准 ASCII 字符集，控制字符的 ASCII 编码介于 0x00 (NUL) 和 0x1f (US) 之间，以及 0x7f (DEL)，某些平台的特定编译器实现还可以在扩展字符集（0x7f 以上）中定义额外的控制字符。

TODO not clear
*/

#include <stdio.h>
#include <ctype.h>

#define MAXLINE 100 /* max number of chars in one line */
#define OCTLEN  6   /* length of octal values          */

int inc(int pos, int n);

/* print aibitrary input in a sensible way */
int main(int argc, char const *argv[])
{
  int c;
  int pos = 0; /* position in the line */

  while ((c = getchar()) != EOF) {
    if (iscntrl(c) || c == ' ') { /* non-graphic or blank */
      pos = inc(pos, OCTLEN);
      printf(" \\%03o ", c);
      if (c == '\n') {
        pos = 0;
        putchar('\n');
      }
    } else { /* graphic character */
      pos = inc(pos, 1);
      putchar(c);
    }
  }

  return 0;
}

/* increment position counter for output */
int inc(int pos, int n) {
  if (pos+n < MAXLINE) {
    return pos + n;
  } else {
    putchar('\n');
    return n;
  }
}
