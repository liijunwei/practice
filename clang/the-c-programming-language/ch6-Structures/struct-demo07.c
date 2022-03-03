/*
page 115

在所有运算符中, 下面4个运算符的优先级最高:
    + 结构运算符"."和"->"
    + 用于函数调用的"()"
    + 用于下标的"[]"
    +
    + 因此, 他们同操作数之间的结合也最紧密

示例:
struct {
  int len;
  char *str;
} *p;

++p->len

将增加len的值, 而不是增加p的值, 因为其中的隐含括号关系为 ++(p->len)

可以使用括号改变结合次序

例如 (++p)->len 将先执行p的加一操作, 再对len进行操作

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
