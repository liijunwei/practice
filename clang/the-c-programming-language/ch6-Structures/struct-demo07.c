/*
page 115

在所有运算符中, 下面4个运算符的优先级最高
结构运算符 '.' 和 '->'
用于函数调用的 '()'
用于下标的 '[]'
因此, 他们与操作时之间的结合也最紧密
*/

#include <stdio.h>

struct demo {
  int len;
  char *str;
};

int main(int argc, char const *argv[])
{
  struct demo c = {1, "hello"};
  struct demo *p = &c;

  printf("demo.len: %d\n", (*p).len);
  printf("p->len: %d\n", p->len);

  ++p->len;    /* 将增加len的值, 而不是增加p的值, 因为其中的隐含括号关系是 ++(p->len) */
  (++p)->len;  /* 先指向为p加1的操作, 再对len执行操作 */
  (p++)->len;  /* 将先对len执行操作, 再将p加1 */
  *p->str;     /* 取的是指针str所指向的对象的值 */
  *p->str++;   /* 先读取指针str所指向的对象的值, 再将str加1 */
  (*p->str)++; /* 将指针str指向的对象值加1 */
  *p++->str;   /* 先读取指针str指向的对象的值, 然后再将p加1 */

  return 0;
}
