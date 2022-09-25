/*
page 148

类似于isupper这样的函数可以通过某种方式实现以达到节省空间或时间的目的;
考虑节省空间或时间的实现方式

*/

#include <stdio.h>

/* 空间利用率较高 */
int f_isupper(char c) { return (c >= 'A' && c <= 'Z') ? 1 : 0; }

/* 时间利用率较高         */
/* 因为没有函数调用的开销 */
/* 但是每次执行都要展开   */
/* 使用宏时, 需要注意由于参数可能被求值两次而带来的潜在问题 */
/* 对于那些可能对参数进行多次求值的宏, 一定要提高警惕; <ctype.h>
 * 中的toupper和tolower是很好的学习例子 */
#define m_isupper(c) ((c) >= 'A' && (c) <= 'Z') ? 1 : 0

int main(int argc, char const *argv[]) {
  printf("f_isupper\n");
  printf("A: %d\n", f_isupper('A'));
  printf("a: %d\n", f_isupper('a'));

  printf("\n");

  printf("m_isupper\n");
  printf("A: %d\n", m_isupper('A'));
  printf("a: %d\n", m_isupper('a'));

  return 0;
}
